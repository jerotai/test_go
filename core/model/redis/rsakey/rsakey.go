package rsakey

import (
	"github.com/go-redis/redis"
	"Stingray/core/model"
	time "time"
	"fmt"
)

type RsaKey struct {
	data_connect *redis.Client
}

func RsaKeyService() (rsaKey *RsaKey) {
	return &RsaKey{}
}

func (d *RsaKey) Init() {
	d.data_connect, _ = model.NewRedisConn().GetRsaKeyMasterConnect()
}

/**
 * set login rsa keys
 */
func (d *RsaKey) SetTokenRsaKey(token string, publicKey string, privateKey string) (string, error) {
	data := make(map[string]interface{})
	data["publicKey"] = publicKey
	data["privateKey"] = privateKey
	
	code := d.data_connect.HMSet("rsa_keys:" + token, data)
	d.data_connect.Expire("rsa_keys:" + token, 24 *time.Hour)
	return code.Result()
}

/**
 * get Rsa Private Key By Token
 */
func (d *RsaKey) GetTokenRsaPrivateKey(token string) string {
	fields, _ := d.data_connect.HMGet("rsa_keys:" + token,  "privateKey").Result()
	return fmt.Sprintf("%v", fields[0])
}

/**
 * get Rsa Public Key By Token
 */
func (d *RsaKey) GetTokenRsaPublicKey(token string) string {
	fields, _ := d.data_connect.HMGet("rsa_keys:" + token,  "publicKey").Result()
	return fmt.Sprintf("%v", fields[0])
}
