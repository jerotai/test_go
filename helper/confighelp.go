package helper

import (
	"github.com/spf13/viper"
	"time"
	"os"
	//"runtime"
	//"path"
	"strings"
	"sync"
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

var (
	// 單例模式實例
	redisInst *viper.Viper
	apiInst *viper.Viper
	redisConfOnce   sync.Once
	apiConfOnce    sync.Once
)

// apiForge : 取得實例
func apiForge() *viper.Viper {
	apiConfOnce.Do(func() {
		apiInst = viper.New()

		apiInst.AddConfigPath(configPath())
		apiInst.SetConfigName("apisetting")
		apiInst.SetEnvPrefix("router")
		apiInst.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		apiInst.AutomaticEnv()
		
		if err := apiInst.ReadInConfig(); err != nil {
			HelperLog.ErrorLog("[Confighelp ApiSetting] " + err.Error())
		}
	})
	
	return apiInst
}

//redisForge: 取得實例
func redisForge() *viper.Viper {
	redisConfOnce.Do(func() {
		redisInst = viper.New()
		redisInst.SetConfigName("redis")
		redisInst.AddConfigPath(configPath())
		redisInst.SetEnvPrefix("router")
		redisInst.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		redisInst.AutomaticEnv()
		
		err := redisInst.ReadInConfig()
		
		if err != nil {
			HelperLog.ErrorLog("[confighelp RedisConf] " + err.Error())
		}
	})
	
	return redisInst
}

/**
 * 取得設定檔路徑
 */
func configPath() string {
	confPath := ""
	if os.Getenv("ROUTER_PATH") == "" {
		/*
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			errStr := "[confighelp configPath] No caller information"
			HelperLog.ErrorLog(errStr)
			panic(errStr)
		}
		confPath = path.Dir(filename) + "/../config"
		*/
		confPath = "./config"
	} else {
		confPath = os.Getenv("ROUTER_PATH") + "/config"
	}
	
	return confPath
}

/**
 * 取得 redis config
 */
func RedisConf(name string, ctype string) *redisConf {
	conf := redisForge()
	
	reConf := &redisConf{}
	reConf.Host = conf.GetString(name + "." + ctype + ".host")
	reConf.Port = conf.GetString(name + "." + ctype + ".port")
	reConf.Password = conf.GetString(name + "." + ctype + ".password")
	reConf.DB = conf.GetInt(name + "." + ctype + ".db")
	
	return reConf
}

/**
 * 取得 redis setting conf
 */
func RedisSettingConf() *redisSettingConf {
	conf := redisForge()
	reConf := &redisSettingConf{}
	
	reConf.DialTimeout = conf.GetDuration("redis_info.dial_timeout")
	reConf.ReadTimeout = conf.GetDuration("redis_info.read_timeout")
	reConf.WriteTimeout = conf.GetDuration("redis_info.write_timeout")
	reConf.PoolSize = conf.GetInt("redis_info.pool_size")
	reConf.PoolTimeout = conf.GetDuration("redis_info.pool_timeout")
	
	return reConf
}

type ApiSettingConf struct {
	Host string
	Port int
}

func ApiSetting(name string) *ApiSettingConf {
	conf := apiForge()
	reConf := &ApiSettingConf{}
	reConf.Host = conf.GetString(name + ".host")
	reConf.Port = conf.GetInt(name+ ".port")
	
	return reConf
}