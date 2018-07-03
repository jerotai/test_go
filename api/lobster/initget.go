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
	default:
		//todo
		helper.HelperLog.ErrorLog("[Lobster InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
