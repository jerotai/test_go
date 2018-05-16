package helper

import (
	"testing"
	
	//"github.com/spf13/viper"
	//"Stingray/helper"
	"Stingray/helper"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)


// 測試 RSA Key
func TestRsaPrivateKey(t *testing.T) {
	viper.SetConfigName("rsa")
	viper.AddConfigPath("../../config/")
	
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	
	vk := viper.GetString("privateKey")
	assert.Equal(t, vk, helper.RsaPrivateKey(), "privateKey error")
}

// 測試 RSA Key
func TestRsaPublicKey(t *testing.T) {
	viper.SetConfigName("rsa")
	viper.AddConfigPath("../../config/")
	
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	
	vk := viper.GetString("publicKey")
	assert.Equal(t, vk, helper.RsaPublicKey(), "publicKey error")
}