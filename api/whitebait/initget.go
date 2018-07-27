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
		apiRequestUrl = "user/alive"
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
	case "/user_level/user_list":
		apiDto = &dto.BackEndUserLevelUserList{}
		apiRequestUrl = "user_level/user_list"
	case "/user_level/userPreview":
		apiDto = &dto.UserLevelUserPreview{}
		apiRequestUrl = "user_level/user/preview"
	case "/user_level/batchPreview":
		apiDto = &dto.UserLevelBatchPreview{}
		apiRequestUrl = "user_level/batch/preview"
	case "/user_level/dropdownList":
		apiDto = &dto.UserLevelDropdownList{}
		apiRequestUrl = "user_level/dropdownlist"
	case "/bank/transfer":
		apiDto = &dto.BankTransger{}
		apiRequestUrl = "bank/transfer"
	case "/bank/third":
		apiDto = &dto.BankThird{}
		apiRequestUrl = "bank/third"
	case "/user/passwordWithdrawCheck":
		apiDto = &dto.UserPasswordWithdrawCheck{}
		apiRequestUrl = "user/password/withdraw/check"
	case "/user/blurry":
		apiDto = &dto.UserBlurry{}
		apiRequestUrl = "user/blurry"
	case "/guest/myself":
		apiDto = &dto.GuestMyself{}
		apiRequestUrl = "guest/myself"
	case "/guest/hall":
		apiDto = &dto.GuestHall{}
		apiRequestUrl = "guest/hall"
	case "/guest/site":
		apiDto = &dto.GuestSite{}
		apiRequestUrl = "guest/site"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Whitebait InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
