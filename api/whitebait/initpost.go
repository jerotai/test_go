package whitebait

import (
	dto "Stingray/core/dto/whitebaitdto"
	"Stingray/helper"
)

/**
 * init Api Post Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Whitebait) InitPostApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/user/register":
		apiDto = &dto.CreateUser{}
		apiRequestUrl = "user/register"
	case "/user/create":
		apiDto = &dto.BackEndCreateUser{}
		apiRequestUrl = "user/create"
	case "/user_level":
		apiDto = &dto.CreateUserLevel{}
		apiRequestUrl = "user_level"
	case "/user/passwordWithdraw":
		apiDto = &dto.CreateUserPasswordWithdraw{}
		apiRequestUrl = "user/password/withdraw"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whitebait InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}