package whitebait

import (
	dto "Stingray/core/dto/whitebaitdto"
	"Stingray/helper"
)

/**
 * init Api Put Config
 * dto = &dto.Menu{}
 * apiRequestUrl = "menu"
 *
 */
func (i *Whitebait) InitPutApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/user/password":
		apiDto = &dto.UpdataUserPassword{}
		apiRequestUrl = "user/password"
	case "/user/info":
		apiDto = &dto.UpdataUserInfo{}
		apiRequestUrl = "user/info"
	case "/user/registerSetting":
		apiDto = &dto.UpdataUserRegisterSetting{}
		apiRequestUrl = "user/register/setting"
	case "/user_level":
		apiDto = &dto.UpdataUserLevel{}
		apiRequestUrl = "user_level"
	case "/user_level/amount":
		apiDto = &dto.UpdataUserLevelAmount{}
		apiRequestUrl = "user_level/amount"
	case "/user_level/payment":
		apiDto = &dto.UpdataUserLevelPayment{}
		apiRequestUrl = "user_level/payment"
	case "/user_level/company_bank":
		apiDto = &dto.UpdataUserLevelCompanyBank{}
		apiRequestUrl = "user_level/company_bank"
	case "/user/passwordWithdraw":
		apiDto = &dto.UpdataUserPasswordWithdraw{}
		apiRequestUrl = "user/password/withdraw"
	case "/user/passwordWithdrawUpdate":
		apiDto = &dto.UpdataBackEndUserPasswordWithdraw{}
		apiRequestUrl = "user/password/withdraw/update"
	case "/user/passwordUpdate":
		apiDto = &dto.UpdataBackEndUserPassword{}
		apiRequestUrl = "user/password/update"
	case "/user_level/user":
		apiDto = &dto.UpdataBackEndUserLevel{}
		apiRequestUrl = "user_level/user"
	case "/user_level/batch":
		apiDto = &dto.UpdataBackEndUserLevelBatch{}
		apiRequestUrl = "user_level/user"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whitebait InitPutApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}