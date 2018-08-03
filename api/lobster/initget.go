package lobster

import (
	dto "Stingray/core/dto/lobsterdto"
	"Stingray/helper"
)

/**
 * init Api Get Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (l *Lobster) InitGetApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/report/outInList":
		apiDto = &dto.ReportOotInList{}
		apiRequestUrl = "report/out_in"
	case "/report/depositList":
		apiDto = &dto.ReportDepositList{}
		apiRequestUrl = "report/deposit"
	case "/report/withdrawList":
		apiDto = &dto.ReportWithdrawList{}
		apiRequestUrl = "report/withdraw"
	case "/report/userTrans":
		apiDto = &dto.ReportUserTrans{}
		apiRequestUrl = "report/user_trans"
	case "/kind/depositType":
		apiDto = &dto.DepositType{}
		apiRequestUrl = "/kind/deposit_type"
	case "/kind/withdrawType":
		apiDto = &dto.WithdrawType{}
		apiRequestUrl = "kind/withdraw_type"
	case "/kind/transType":
		apiDto = &dto.TransType{}
		apiRequestUrl = "kind/trans_type"
	case "/userLoginReport/userLoginDriveCount":
		apiDto = &dto.UserLoginDriveCount{}
		apiRequestUrl = "userLoginReport/user_login_drive_count"
	case "/userLoginReport/userLoginInfo":
		apiDto = &dto.UserLoginInfo{}
		apiRequestUrl = "userLoginReport/user_login_info"
	case "/userLoginReport/userLoginInfoOnline":
		apiDto = &dto.UserLoginInfoOnline{}
		apiRequestUrl = "userLoginReport/user_login_online"
	case "/userLoginReport/userLoginRecord":
		apiDto = &dto.UserLoginRecord{}
		apiRequestUrl = "userLoginReport/user_login_record"
	case "/betting/subordinate":
		apiDto = &dto.BettingSubordinate{}
		apiRequestUrl = "betting/subordinate"
	case "/betting/subordinateSummary":
		apiDto = &dto.BettingSubordinateSummary{}
		apiRequestUrl = "betting/subordinate/summary"
	case "/betting/game":
		apiDto = &dto.BettingGame{}
		apiRequestUrl = "betting/game"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Lobster InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
