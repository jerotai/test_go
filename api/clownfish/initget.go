package clownfish


import (
	"Stingray/helper"
	dto "Stingray/core/dto/clownfishdto"
)

/**
 * init Api Get Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Clownfish) InitGetApiConfig(apiUrl string) (string, interface{}){
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/cooperation/fortune":
		apiDto = &dto.CooperationFortune{}
		apiRequestUrl = "cooperation/fortune"
	case "/cooperation/fortuneUser":
		apiDto = &dto.CooperationFortuneUser{}
		apiRequestUrl = "cooperation/fortune/user"
	case "game/providerList":
		apiDto = &dto.GameProviderList{}
		apiRequestUrl = "game/provider/list"
	case "game/providerKind":
		apiDto = &dto.GameProviderKind{}
		apiRequestUrl = "game/provider/kind"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Clownfish InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
