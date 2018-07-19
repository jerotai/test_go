package jellyfish

import (
	dto "Stingray/core/dto/jellyfishdto"
	"Stingray/helper"
)

/**
 * init Api Post Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (j *JellyFish) InitPostApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/login":
		apiDto = &dto.Login{}
		apiRequestUrl = "login"
	case "/hall":
		apiDto = &dto.CreateHall{}
		apiRequestUrl = "hall"
	case "/shareholder":
		apiDto = &dto.CreateShareholder{}
		apiRequestUrl = "shareholder"
	case "/shareholderSub":
		apiDto = &dto.CreateShareholderSub{}
		apiRequestUrl = "shareholderSub"
	case "/hallSub":
		apiDto = &dto.CreateHallSub{}
		apiRequestUrl = "hallSub"
	case "/authGroup":
		apiDto = &dto.CreateAuthGroup{}
		apiRequestUrl = "authGroup"
	case "/generalAgent":
		apiDto = &dto.CreateGeneralAgent{}
		apiRequestUrl = "generalAgent"
	case "/generalAgentSub":
		apiDto = &dto.CreateGeneralAgentSub{}
		apiRequestUrl = "generalAgentSub"
	case "/agent":
		apiDto = &dto.CreateAgent{}
		apiRequestUrl = "agent"
	case "/agentSub":
		apiDto = &dto.CreateAgentSub{}
		apiRequestUrl = "agentSub"
	default:
		//todo
		helper.HelperLog.ErrorLog("[JellyFish InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}