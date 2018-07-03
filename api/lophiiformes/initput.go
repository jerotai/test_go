package lophiiformes

import (
	dto "Stingray/core/dto/lophiiformesdto"
	
	"Stingray/helper"
)

/**
 * init Api Put Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (l *Lophiiformes) InitPutApiConfig(apiUrl string) (string, interface{})  {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/company_bank":
		apiDto = &dto.UpdataCompanyBank{}
		apiRequestUrl = "company_bank"
	case "/payment_config":
		apiDto = &dto.UpdataPaymentConfig{}
		apiRequestUrl = "payment_config"
	case "/payment_config/banks":
		apiDto = &dto.UpdataPaymentConfigBanks{}
		apiRequestUrl = "payment_config/banks"
	case "/user_bank":
		apiDto = &dto.UpdataUserBank{}
		apiRequestUrl = "user_bank"
	case "/user_bank/backend":
		apiDto = &dto.BackEndUpdataUserBank{}
		apiRequestUrl = "user_bank/backend"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Lophiiformes InitPutApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}