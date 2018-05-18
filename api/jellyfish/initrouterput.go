package jellyfish

import (
	"fmt"
	dto "Stingray/core/dto/jellyfishdto"
)

/**
 * init Api Put Config
 */
func (j *JellyFish) InitPutApiConfig(apiUrl string) {
	fmt.Println("Put", apiUrl)
	switch apiUrl {
	case "/hall":
		j.ApiConf.dto = &dto.UpdataHall{}
		j.ApiConf.apiRequestUrl = "hall"
	case "/hall/enabled":
		j.ApiConf.dto = &dto.EnabledHall{}
		j.ApiConf.apiRequestUrl = "hall/enabled"
	case "/shareholder":
		j.ApiConf.dto = &dto.UpdataShareholder{}
		j.ApiConf.apiRequestUrl = "shareholder"
	case "/hallSub":
		j.ApiConf.dto = &dto.UpdataShareholder{}
		j.ApiConf.apiRequestUrl = "hallSub"
	case "/authGroup":
		j.ApiConf.dto = &dto.UpdataAuthGroup{}
		j.ApiConf.apiRequestUrl = "authGroup"
	default:
		//todo
	}
}