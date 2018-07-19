package trait

import (
	"Stingray/core/model/redis/rsakey"
	"Stingray/helper"
)

func UpdateRsaKeyExpire(ApiToken string) {
	rsaKeyRedis := rsakey.New()
	UpdateStatus := rsaKeyRedis.SetTokenRsaKeyExpire(ApiToken)
	
	if UpdateStatus != true {
		helper.HelperLog.ErrorLog("[UpdateRsaKeyExpire] Update Expire Error")
	}
}

/**
 * save login rsa keys to redis
 */
func SaveLoginRsaKey(token string, rsaPublicKey string, rsaPrivateKey string) {
	rsaKeyRedis := rsakey.New()
	status, err :=  rsaKeyRedis.SetTokenRsaKey(token, rsaPublicKey, rsaPrivateKey)
	
	if (err != nil || status != "OK") {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[trait-SaveLoginRsaKey] " + err.Error())
		panic(err)
	}
}
