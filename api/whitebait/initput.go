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
		apiDto = &dto.UpdateUserPassword{}
		apiRequestUrl = "user/password"
	case "/user/info":
		apiDto = &dto.UpdateUserInfo{}
		apiRequestUrl = "user/info"
	case "/user/registerSetting":
		apiDto = &dto.UpdateUserRegisterSetting{}
		apiRequestUrl = "user/register/setting"
	case "/user_level":
		apiDto = &dto.UpdateUserLevel{}
		apiRequestUrl = "user_level"
	case "/user_level/amount":
		apiDto = &dto.UpdateUserLevelAmount{}
		apiRequestUrl = "user_level/amount"
	case "/user_level/payment":
		apiDto = &dto.UpdateUserLevelPayment{}
		apiRequestUrl = "user_level/payment"
	case "/user_level/company_bank":
		apiDto = &dto.UpdateUserLevelCompanyBank{}
		apiRequestUrl = "user_level/company_bank"
	case "/user/passwordWithdraw":
		apiDto = &dto.UpdateUserPasswordWithdraw{}
		apiRequestUrl = "user/password/withdraw"
	case "/user/passwordWithdrawUpdate":
		apiDto = &dto.UpdateBackEndUserPasswordWithdraw{}
		apiRequestUrl = "user/password/withdraw/update"
	case "/user/passwordUpdate":
		apiDto = &dto.UpdateBackEndUserPassword{}
		apiRequestUrl = "user/password/update"
	case "/user_level/user":
		apiDto = &dto.UpdateBackEndUserLevel{}
		apiRequestUrl = "user_level/user"
	case "/user_level/batch":
		apiDto = &dto.UpdateBackEndUserLevelBatch{}
		apiRequestUrl = "user_level/batch"
	case "/guest/myself":
		apiDto = &dto.UpdateGuestMyself{}
		apiRequestUrl = "guest/myself"
	case "/guest/hall":
		apiDto = &dto.UpdateGuestHall{}
		apiRequestUrl = "guest/hall"
	case "/guest/site":
		apiDto = &dto.UpdateGuestSite{}
		apiRequestUrl = "guest/site"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whitebait InitPutApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}