package model

import (
	"github.com/go-redis/redis"
	"routes/helper"
	"time"
	"fmt"
)

type RedisConnect struct {
}

func NeqwRedisConn() (*RedisConnect) {
	return &RedisConnect{}
}

func (r *RedisConnect) Test() {
	conf := helper.RedisConf("site_list", "master")
	redisInfoConf := helper.RedisSettingConf()
	
	fmt.Println(conf)
	fmt.Println(redisInfoConf)
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
 * 取得 redis connect
 */
func (r *RedisConnect) connect(name string, ctype string) (*redis.Client, error) {
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
		// log
	}
	
	return client, err
}