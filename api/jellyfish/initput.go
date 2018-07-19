package jellyfish

import (
	dto "Stingray/core/dto/jellyfishdto"
	"Stingray/helper"
)

/**
 * init Api Put Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (j *JellyFish) InitPutApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	
	switch apiUrl {
	case "/password":
		apiDto = &dto.UpdatePassWord{}
		apiRequestUrl = "password"
	case "/hall":
		apiDto = &dto.UpdateHall{}
		apiRequestUrl = "hall"
	case "/hall/password":
		apiDto = &dto.UpdateHallPassword{}
		apiRequestUrl = "hall/password"
	case "/hall/enabled":
		apiDto = &dto.EnabledHall{}
		apiRequestUrl = "hall/enabled"
	case "/shareholder":
		apiDto = &dto.UpdateShareholder{}
		apiRequestUrl = "shareholder"
	case "/shareholder/password":
		apiDto = &dto.UpdateShareholderPassword{}
		apiRequestUrl = "shareholder/password"
	case "/shareholderSub":
		apiDto = &dto.UpdateShareholderSub{}
		apiRequestUrl = "shareholderSub"
	case "/shareholderSub/password":
		apiDto = &dto.UpdateShareholderSubPassword{}
		apiRequestUrl = "shareholderSub/password"
	case "/hallSub":
		apiDto = &dto.UpdateHallSub{}
		apiRequestUrl = "hallSub"
	case "/hallSub/password":
		apiDto = &dto.UpdateHallSubPassword{}
		apiRequestUrl = "hallSub/password"
	case "/authGroup":
		apiDto = &dto.UpdateAuthGroup{}
		apiRequestUrl = "authGroup"
	case "/generalAgent":
		apiDto = &dto.UpdateGeneralAgent{}
		apiRequestUrl = "generalAgent"
	case "/generalAgent/password":
		apiDto = &dto.UpdateGeneralAgentPassword{}
		apiRequestUrl = "generalAgent/password"
	case "/generalAgentSub":
		apiDto = &dto.UpdateGeneralAgentSub{}
		apiRequestUrl = "generalAgentSub"
	case "/generalAgentSub/password":
		apiDto = &dto.UpdateGeneralAgentSubPassword{}
		apiRequestUrl = "generalAgentSub/password"
	case "/agent":
		apiDto = &dto.UpdateAgent{}
		apiRequestUrl = "agent"
	case "/agent/password":
		apiDto = &dto.UpdateAgentPassword{}
		apiRequestUrl = "agent/password"
	case "/agentSub":
		apiDto = &dto.UpdateAgentSub{}
		apiRequestUrl = "agentSub"
	case "/agentSub/password":
		apiDto = &dto.UpdateAgentSubPassword{}
		apiRequestUrl = "agentSub/password"
	default:
		//todo
		helper.HelperLog.ErrorLog("[JellyFish InitPurApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}