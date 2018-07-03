package whitebait

import (
	dto "Stingray/core/dto/whitebaitdto"
	"Stingray/helper"
)

/**
 * init Api Get Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Whitebait) InitGetApiConfig(apiUrl string) (string, interface{}){
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/user/list":
		apiDto = &dto.UserList{}
		apiRequestUrl = "user/list"
	case "/user/register":
		apiDto = &dto.UserRegister{}
		apiRequestUrl = "user/register"
	case "/user/promo_code":
		apiDto = &dto.UserPromoCode{}
		apiRequestUrl = "user/promo_code"
	case "/user/plat":
		apiDto = &dto.UserPlat{}
		apiRequestUrl = "user/plat"
	case "/user/data":
		apiDto = &dto.UserData{}
		apiRequestUrl = "user/data"
	case "/user/balance":
		apiDto = &dto.UserBalance{}
		apiRequestUrl = "user/balance"
	case "/user/company_bank":
		apiDto = &dto.UserCompanyBank{}
		apiRequestUrl = "user/company_bank"
	case "/user/info":
		apiDto = &dto.UserInfo{}
		apiRequestUrl = "user/info"
	case "/user/registerSetting":
		apiDto = &dto.UserRegisterSetting{}
		apiRequestUrl = "user/register/setting"
	case "/user/alive":
		apiDto = &dto.UserAlive{}
		apiRequestUrl = "user/register/alive"
	case "/user_level/list":
		apiDto = &dto.UserLevel{}
		apiRequestUrl = "user_level"
	case "/user_level/data":
		apiDto = &dto.UserLevelData{}
		apiRequestUrl = "user_level/data"
	case "/user_level/amount":
		apiDto = &dto.UserLevelAmount{}
		apiRequestUrl = "user_level/amount"
	case "/user_level/payment":
		apiDto = &dto.UserLevelPayment{}
		apiRequestUrl = "user_level/payment"
	case "/user_level/company_bank":
		apiDto = &dto.UserLevelCompanyBankList{}
		apiRequestUrl = "user_level/company_bank"
	case "/bank/transfer":
		apiDto = &dto.BankTransger{}
		apiRequestUrl = "bank/transfer"
	case "/bank/third":
		apiDto = &dto.BankThird{}
		apiRequestUrl = "bank/third"
	case "/user/passwordWithdrawCheck":
		apiDto = &dto.UserPasswordWithdrawCheck{}
		apiRequestUrl = "user/password/withdraw/check"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whitebait InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
