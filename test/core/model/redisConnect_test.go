package model

import (
	"routes/core/model"
	
	"testing"
	"fmt"
)

/**
 * 測試 Site Master Connect
 */
func TestGetSiteListMasterConnect(t *testing.T) {
	redisConnect, err := model.NewRedisConn().GetSiteListMasterConnect()
	
	if err != nil {
		fmt.Println(err)
		//錯誤處理
	}
	defer redisConnect.Close()
}

/**
 * 測試 Site Slave Connect
 */
func TestGetSiteListSlaveConnect(t *testing.T) {
	redisConnect, err := model.NewRedisConn().GetSiteListSlaveConnect()
	
	if err != nil {
		fmt.Println(err)
		//錯誤處理
	}
	defer redisConnect.Close()
}