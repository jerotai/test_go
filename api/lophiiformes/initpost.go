package lophiiformes

import (
	dto "Stingray/core/dto/lophiiformesdto"
	
	"Stingray/helper"
)

/**
 * init Api Post Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (l *Lophiiformes) InitPostApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/company_bank":
		apiDto = &dto.CreateCompanyBank{}
		apiRequestUrl = "company_bank"
	case "/payment_config":
		apiDto = &dto.CreatePaymentConfig{}
		apiRequestUrl = "payment_config"
	case "/user_bank":
		apiDto = &dto.CreateUserBank{}
		apiRequestUrl = "user_bank"
	case "/user_bank/backend":
		apiDto = &dto.BackEndCreateUserBank{}
		apiRequestUrl = "user_bank/backend"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Lophiiformes InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}