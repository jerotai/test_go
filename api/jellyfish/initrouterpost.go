package jellyfish

import (
	"fmt"
	dto "Stingray/core/dto/jellyfishdto"
)

/**
 * init Api Post Config
 */
func (j *JellyFish) InitPostApiConfig(apiUrl string) {
	fmt.Println("Post", apiUrl)
	switch apiUrl {
	case "/hall":
		j.ApiConf.dto = &dto.Hall{}
		j.ApiConf.apiRequestUrl = "hall"
	case "/shareholder":
		j.ApiConf.dto = &dto.CreateShareholder{}
		j.ApiConf.apiRequestUrl = "shareholder"
	case "/hallSub":
		j.ApiConf.dto = &dto.CreateHallSub{}
		j.ApiConf.apiRequestUrl = "hallSub"
	case "/authGroup":
		j.ApiConf.dto = &dto.CreateAuthGroup{}
		j.ApiConf.apiRequestUrl = "authGroup"
	default:
		//todo
	}
}