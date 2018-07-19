package octopus

import (
	dto "Stingray/core/dto/octopusdto"
	"Stingray/helper"
)

/**
 * init Api Get Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Octopus) InitGetApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/manualDeposits/list":
		apiDto = &dto.ManualDepositsList{}
		apiRequestUrl = "manualDeposits"
	case "/manualDeposits/audit":
		apiDto = &dto.ManualDepositsAudit{}
		apiRequestUrl = "manualDeposits/audit"
	case "/manualWithdraws/list":
		apiDto = &dto.ManualWithdrawsList{}
		apiRequestUrl = "manualWithdraws"
	case "/manualWithdraws/audit":
		apiDto = &dto.ManualWithdrawsAudit{}
		apiRequestUrl = "manualWithdraws/audit"
	case "/manualWithdraws/grant":
		apiDto = &dto.ManualWithdrawsGrant{}
		apiRequestUrl = "manualWithdraws/grant"
	case "/fourthDeposits/list":
		apiDto = &dto.FourthDepositsList{}
		apiRequestUrl = "fourthDeposits"
	case "/fourthDeposits/audit":
		apiDto = &dto.FourthDepositsAudit{}
		apiRequestUrl = "fourthDeposits/audit"
	case "/fourthDeposits/limit":
		apiDto = &dto.FourthDepositsLimit{}
		apiRequestUrl = "fourthDeposits/limit"
	case "/bankDeposits/list":
		apiDto = &dto.BankList{}
		apiRequestUrl = "bankDeposits"
	case "/bankDeposits/menu":
		apiDto = &dto.BankDepositsMenu{}
		apiRequestUrl = "bankDeposits/menu"
	case "/providerDeposits/list":
		apiDto = &dto.ProviderDepositsList{}
		apiRequestUrl = "providerDeposits"
	case "/providerDeposits/audit":
		apiDto = &dto.ProviderDepositsAudit{}
		apiRequestUrl = "providerDeposits/audit"
	case "/providerDeposits/limit":
		apiDto = &dto.ProviderDepositsLimit{}
		apiRequestUrl = "providerDeposits/limit"
	case "/providerWithdraws/list":
		apiDto = &dto.ProviderWithdrawsList{}
		apiRequestUrl = "providerWithdraws"
	case "/providerWithdraws/limit":
		apiDto = &dto.ProviderWithdrawsLimit{}
		apiRequestUrl = "providerWithdraws/limit"
	case "/providerWithdraws/audit":
		apiDto = &dto.ProviderWithdrawsAudit{}
		apiRequestUrl = "providerWithdraws/audit"
	case "/providerWithdraws/grant":
		apiDto = &dto.ProviderWithdrawsGrant{}
		apiRequestUrl = "providerWithdraws/grant"
	case "/fourthDeposits/menu_thirds":
		apiDto = &dto.FourthDepositsMenuThirds{}
		apiRequestUrl = "fourthDeposits/menu/thirds"
	case "/fourthDeposits/menu_fourths":
		apiDto = &dto.FourthDepositsMenuFourths{}
		apiRequestUrl = "fourthDeposits/menu/fourths"
	case "/fourthDeposits/menu":
		apiDto = &dto.FourthDepositsMenu{}
		apiRequestUrl = "fourthDeposits/menu"
	case "/bankDeposits/front_menu":
		apiDto = &dto.BankDepositsFrontMenu{}
		apiRequestUrl = "bankDeposits/front_menu"
	case "/bankDeposits/audit":
		apiDto = &dto.BankDepositsAudit{}
		apiRequestUrl = "bankDeposits/audit"
	case "/transReports":
		apiDto = &dto.TransReportList{}
		apiRequestUrl = "transReports"
	case "/transReports/deposit":
		apiDto = &dto.TransReportDeposit{}
		apiRequestUrl = "transReports/deposit"
	case "/transReports/withdrawList":
		apiDto = &dto.TransReportWithdraw{}
		apiRequestUrl = "transReports/withdraw"
	case "/transSumReport/depositList":
		apiDto = &dto.TransSumReportDepositList{}
		apiRequestUrl = "transSumReport/deposit"
	case "/transSumReport/withdrawList":
		apiDto = &dto.TransSumReportWithdrawList{}
		apiRequestUrl = "transSumReport/withdraw"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Octopus InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
