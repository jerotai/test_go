package model

import (
	"github.com/go-redis/redis"
	"Stingray/helper"
	"time"
	"fmt"
	"sync"
)

type RedisConnect struct {
}

var (
	//Redis 連線物件
	pool map[string]*redis.Client
	
	//同步鎖
	mu *sync.Mutex
)

func init() {
	pool = make(map[string]*redis.Client)
	mu = &sync.Mutex{}
}

func NewRedisConn() (*RedisConnect) {
	return &RedisConnect{}
}

/**
 * 取得 Site Info Master Connect
 */
func (r *RedisConnect) GetSiteListMasterConnect() (*redis.Client, error) {
	return r.connect("site_list", "master")
}

/**
 * 取得 Site Info Slave Connect
 */
func (r *RedisConnect) GetSiteListSlaveConnect() (*redis.Client, error) {
	return r.connect("site_list", "slave")
}

/**
 * 取得 Rsa Key Master Connect
 */
func (r *RedisConnect) GetRsaKeyMasterConnect() (*redis.Client, error) {
	return r.connect("rsa_key", "master")
}

/**
 * 取得 Rsa Key Slave Connect
 */
func (r *RedisConnect) GetRsaKeySlaveConnect() (*redis.Client, error) {
	return r.connect("rsa_key", "slave")
}

/**
 * 取得 redis connect
 */
func (r *RedisConnect) connect(name string, ctype string) (*redis.Client, error) {
	mu.Lock()
	defer mu.Unlock()
	
	if conn, ok := pool[name]; ok {
		return conn, nil
	}
	
	conf := helper.RedisConf(name, ctype)
	redisInfoConf := helper.RedisSettingConf()
	
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password, // no password set
		DB:       conf.DB,  // use default DB
		DialTimeout:  redisInfoConf.DialTimeout * time.Second,
		ReadTimeout:  redisInfoConf.ReadTimeout * time.Second,
		WriteTimeout: redisInfoConf.WriteTimeout * time.Second,
		PoolSize:     redisInfoConf.PoolSize,
		PoolTimeout:  redisInfoConf.PoolTimeout * time.Second,
	})
	
	_, err := client.Ping().Result()
	
	if err != nil {
		helper.HelperLog.ErrorLog(fmt.Sprintf("Redis Connect Error : %s", err.Error()))
	}
	
	pool[name] = client
	
	return client, err
}