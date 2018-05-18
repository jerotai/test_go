package jellyfish

import (
	"fmt"
	dto "Stingray/core/dto/jellyfishdto"
)

/**
 * init Api Get Config
 */
func (j *JellyFish) InitGetApiConfig(apiUrl string) {
	fmt.Println("Get", apiUrl)
	switch apiUrl {
	case "/hall/subList":
		j.ApiConf.dto = &dto.SubList{}
		j.ApiConf.apiRequestUrl = "hall/subList"
	case "/shareholder/list":
		j.ApiConf.dto = &dto.ShareholderList{}
		j.ApiConf.apiRequestUrl = "shareholder"
	case "/shareholder/data":
		j.ApiConf.dto = &dto.ShareholderData{}
		j.ApiConf.apiRequestUrl = "shareholder/data"
	case "/hallSub/list":
		j.ApiConf.dto = &dto.ShareholderList{}
		j.ApiConf.apiRequestUrl = "hallSub"
	case "/authGroup/list":
		j.ApiConf.dto = &dto.AuthGroupList{}
		j.ApiConf.apiRequestUrl = "authGroup"
	case "/authGroup/detail":
		j.ApiConf.dto = &dto.AuthGroupDatail{}
		j.ApiConf.apiRequestUrl = "authGroup/detail"
	case "/authGroup/dropDown":
		j.ApiConf.dto = &dto.AuthGroupDropDown{}
		j.ApiConf.apiRequestUrl = "authGroup/dropDown"
	default:
		fmt.Println("Get Error")
		//todo
	}
}
