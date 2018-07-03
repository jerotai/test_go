package rsakey

import (
	"github.com/go-redis/redis"
	"Stingray/core/model"
	"time"
	"fmt"
	"errors"
	"Stingray/helper"
)

type RsaKey struct {
	data_connect *redis.Client
}

const ExpireTime = 3 //hour

func New() *RsaKey {
	reRsa := &RsaKey{}
	err := errors.New("")
	
	reRsa.data_connect, err = model.NewRedisConn().GetRsaKeyMasterConnect()
	
	if err != nil {
		helper.HelperLog.ErrorLog("[RsaKey New] data_connect Error :" + err.Error())
	}
	
	return reRsa
}

/**
 * set login rsa keys
 */
func (d *RsaKey) SetTokenRsaKey(token string, publicKey string, privateKey string) (string, error) {
	data := make(map[string]interface{})
	data["publicKey"] = publicKey
	data["privateKey"] = privateKey
	
	code := d.data_connect.HMSet("rsa_keys:" + token, data)
	d.data_connect.Expire("rsa_keys:" + token, ExpireTime *time.Hour)
	return code.Result()
}

/**
 * set login rsa keys Expire
 */
func (d *RsaKey) SetTokenRsaKeyExpire(token string) (bool) {
	code := d.data_connect.Expire("rsa_keys:" + token, ExpireTime *time.Hour)
	return code.Val()
}

/**
 * get Rsa Private Key By Token
 */
func (d *RsaKey) GetTokenRsaPrivateKey(token string) string {
	fields, _ := d.data_connect.HMGet("rsa_keys:" + token,  "privateKey").Result()
	
	if fields[0] == nil {
		return ""
	}
	
	return fmt.Sprintf("%v", fields[0])
}

/**
 * get Rsa Public Key By Token
 */
func (d *RsaKey) GetTokenRsaPublicKey(token string) string {
	fields, _ := d.data_connect.HMGet("rsa_keys:" + token,  "publicKey").Result()
	
	if fields[0] == nil {
		return ""
	}
	
	return fmt.Sprintf("%v", fields[0])
}
