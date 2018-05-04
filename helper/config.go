package helper

import (
	"github.com/spf13/viper"
)

func RsaPrivateKey() string {
	viper.SetConfigName("rsa")
	viper.AddConfigPath("../../config/")
	
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	return viper.GetString("privateKey")
}


func RsaPublicKey() string {
	viper.SetConfigName("rsa")
	viper.AddConfigPath("../../config/")
	
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	
	return viper.GetString("publicKey")
}