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
		apiDto = &dto.UpdataPassWord{}
		apiRequestUrl = "password"
	case "/hall":
		apiDto = &dto.UpdataHall{}
		apiRequestUrl = "hall"
	case "/hall/password":
		apiDto = &dto.UpdataHallPassword{}
		apiRequestUrl = "hall/password"
	case "/hall/enabled":
		apiDto = &dto.EnabledHall{}
		apiRequestUrl = "hall/enabled"
	case "/shareholder":
		apiDto = &dto.UpdataShareholder{}
		apiRequestUrl = "shareholder"
	case "/shareholder/password":
		apiDto = &dto.UpdataShareholderPassword{}
		apiRequestUrl = "shareholder/password"
	case "/shareholderSub":
		apiDto = &dto.UpdataShareholderSub{}
		apiRequestUrl = "shareholderSub"
	case "/shareholderSub/password":
		apiDto = &dto.UpdataShareholderSubPassword{}
		apiRequestUrl = "shareholderSub/password"
	case "/hallSub":
		apiDto = &dto.UpdataHallSub{}
		apiRequestUrl = "hallSub"
	case "/hallSub/password":
		apiDto = &dto.UpdataHallSubPassword{}
		apiRequestUrl = "hallSub/password"
	case "/authGroup":
		apiDto = &dto.UpdataAuthGroup{}
		apiRequestUrl = "authGroup"
	case "/generalAgent":
		apiDto = &dto.UpdataGeneralAgent{}
		apiRequestUrl = "generalAgent"
	case "/generalAgent/password":
		apiDto = &dto.UpdataGeneralAgentPassword{}
		apiRequestUrl = "generalAgent/password"
	case "/generalAgentSub":
		apiDto = &dto.UpdataGeneralAgentSub{}
		apiRequestUrl = "generalAgentSub"
	case "/generalAgentSub/password":
		apiDto = &dto.UpdataGeneralAgentSubPassword{}
		apiRequestUrl = "generalAgentSub/password"
	case "/agent":
		apiDto = &dto.UpdataAgent{}
		apiRequestUrl = "agent"
	case "/agent/password":
		apiDto = &dto.UpdataAgentPassword{}
		apiRequestUrl = "agent/password"
	case "/agentSub":
		apiDto = &dto.UpdataAgentSub{}
		apiRequestUrl = "agentSub"
	case "/agentSub/password":
		apiDto = &dto.UpdataAgentSubPassword{}
		apiRequestUrl = "agentSub/password"
	default:
		//todo
		helper.HelperLog.ErrorLog("[JellyFish InitPurApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}