package helper

import (
	"github.com/spf13/viper"
	"fmt"
	"time"
	"path"
	"runtime"
)

type redisConf struct {
	Host string
	Port string
	Password string
	DB int
}

type redisSettingConf struct {
	DialTimeout time.Duration
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PoolSize int
	PoolTimeout time.Duration
}

func helperPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	
	return path.Dir(filename)
}

/**
 * 取得 redis config
 */
func RedisConf(name string, ctype string) *redisConf {
	viper.SetConfigName("redis")
	viper.AddConfigPath(helperPath() + "/../config/")
	err := viper.ReadInConfig()
	
	if err != nil {
		fmt.Printf("%v", err)
	}
	
	subv := viper.Sub(name + "." + ctype)
	
	reConf := &redisConf{}
	
	reConf.Host = subv.GetString("host")
	reConf.Port = subv.GetString("port")
	reConf.Password = subv.GetString("password")
	reConf.DB = subv.GetInt("db")
	
	return reConf
}

/**
 * 取得 redis setting conf
 */
func RedisSettingConf() *redisSettingConf {
	viper.SetConfigName("redis")
	viper.AddConfigPath(helperPath() + "/../config/")
	err := viper.ReadInConfig()
	
	if err != nil {
		fmt.Printf("%v", err)
	}
	
	subv := viper.Sub("redis_info")
	
	reConf := &redisSettingConf{}
	
	reConf.DialTimeout = subv.GetDuration("dial_timeout")
	reConf.ReadTimeout = subv.GetDuration("read_timeout")
	reConf.WriteTimeout = subv.GetDuration("write_timeout")
	reConf.PoolSize = subv.GetInt("pool_size")
	reConf.PoolTimeout = subv.GetDuration("pool_timeout")
	
	return reConf
}

/**
 * 取得RSA私鑰
 */
func RsaPrivateKey() string {
	viper.SetConfigName("rsa")
	viper.AddConfigPath(helperPath() + "/../config/")
	
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	
	return viper.GetString("privateKey")
}

/**
 * 取得RSA公鑰
 */
func RsaPublicKey() string {
	viper.SetConfigName("rsa")
	viper.AddConfigPath(helperPath() + "/../config/")
	
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	
	return viper.GetString("publicKey")
}

type ApiSettingConf struct {
	Host string
	Port string
}

func ApiSetting(name string) *ApiSettingConf {
	viper.SetConfigName("apisetting")
	viper.AddConfigPath(helperPath() + "/../config/")
	
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	
	subv := viper.Sub(name)
	
	reConf := &ApiSettingConf{}
	
	reConf.Host = subv.GetString("Host")
	reConf.Port = subv.GetString("Port")
	
	return reConf
}