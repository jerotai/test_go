package lophiiformes

import (
	dto "Stingray/core/dto/lophiiformesdto"
	
	"Stingray/helper"
)

/**
 * init Api Delete Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (l *Lophiiformes) InitDeleteApiConfig(apiUrl string) (string, interface{}){
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/company_bank":
		apiDto = &dto.DeleteCompanyBank{}
		apiRequestUrl = "company_bank"
	case "/payment_config":
		apiDto = &dto.DeletePaymentConfig{}
		apiRequestUrl = "payment_config"
	case "/user_bank":
		apiDto = &dto.DeleteUserBank{}
		apiRequestUrl = "user_bank"
	case "/user_bank/backend":
		apiDto = &dto.BackEndDeleteUserBank{}
		apiRequestUrl = "user_bank/backend"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Lophiiformes InitDeleteApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}