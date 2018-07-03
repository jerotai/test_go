package helper

import (
	"testing"
	"fmt"
	"Stingray/helper"
)

func TestRedisConf(t *testing.T) {
	redisSiteMaster := helper.RedisConf("site_list", "master")
	if redisSiteMaster.Host == "" || redisSiteMaster.Port == "" {
		t.Error("pRedisConf] site_list master Error")
	}
	
	redisSiteSlave := helper.RedisConf("site_list", "slave")
	
	if redisSiteSlave.Host == "" || redisSiteSlave.Port == "" {
		t.Error("[RedisConf] site_list slave Error")
	}
}

func TestRedisSettingConf(t *testing.T) {
	redisSetting := helper.RedisSettingConf()
	if fmt.Sprintf("%s", redisSetting.PoolSize) == "" ||
			redisSetting.PoolTimeout.String() == "" || redisSetting.ReadTimeout.String() == "" ||
			redisSetting.WriteTimeout.String() == "" || redisSetting.DialTimeout.String() == "" {
		t.Error("RedisSettingConf Error")
	}
}

func TestApiServiceSetting(t *testing.T) {
	var apiList []string
	apiList = append(apiList, "jellyfish_service")
	apiList = append(apiList, "whale_service")
	apiList = append(apiList, "whitebait_service")
	apiList = append(apiList, "lophiiformes_service")
	apiList = append(apiList, "octopus_service")

	for _, v := range apiList {
		apiConf := helper.ApiSetting(v)
		if fmt.Sprintf("%s", apiConf.Port) == "" || apiConf.Host == "" {
			t.Error("[ApiServiceSetting] " + v + "  Error")
		}
	}
}