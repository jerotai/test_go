package octopus

import (
	dto "Stingray/core/dto/octopusdto"
	"Stingray/helper"
)

/**
 * init Api Post Config
 * apiDto = &dto.Menu{}
 * apiRequestUrl = "menu"
 *
 */
func (i *Octopus) InitPostApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/manualDeposits":
		apiDto = &dto.CreateManualDeposits{}
		apiRequestUrl = "manualDeposits"
	case "/manualDeposits/pass":
		apiDto = &dto.ManualDepositsPass{}
		apiRequestUrl = "manualDeposits/pass"
	case "/manualDeposits/reject":
		apiDto = &dto.ManualDepositsReject{}
		apiRequestUrl = "manualDeposits/reject"
	case "/fourthDeposits/pass":
		apiDto = &dto.FourthDepositsPass{}
		apiRequestUrl = "fourthDeposits/pass"
	case "/fourthDeposits":
		apiDto = &dto.CreateFourthDeposits{}
		apiRequestUrl = "fourthDeposits"
	case "/fourthDeposits/reject":
		apiDto = &dto.FourthDepositsReject{}
		apiRequestUrl = "fourthDeposits/reject"
	case "/bankDeposits/pass":
		apiDto = &dto.BankDepositsPass{}
		apiRequestUrl = "bankDeposits/pass"
	case "/bankDeposits":
		apiDto = &dto.CreateBankDeposits{}
		apiRequestUrl = "bankDeposits"
	case "/bankDeposits/reject":
		apiDto = &dto.BankDepositsReject{}
		apiRequestUrl = "bankDeposits/reject"
	case "/providerDeposits":
		apiDto = &dto.CreateProviderDeposits{}
		apiRequestUrl = "providerDeposits"
	case "/providerDeposits/pass":
		apiDto = &dto.ProviderDepositsPass{}
		apiRequestUrl = "providerDeposits/pass"
	case "/providerDeposits/reject":
		apiDto = &dto.ProviderDepositsReject{}
		apiRequestUrl = "providerDeposits/reject"
	case "/manualWithdraws":
		apiDto = &dto.CreateManualWithdraws{}
		apiRequestUrl = "manualWithdraws"
	case "/manualWithdraws/auditPass":
		apiDto = &dto.ManualWithdrawsAuditPass{}
		apiRequestUrl = "manualWithdraws/audit/pass"
	case "/manualWithdraws/auditReject":
		apiDto = &dto.ManualWithdrawsAuditReject{}
		apiRequestUrl = "manualWithdraws/audit/reject"
	case "/manualWithdraws/grantPass":
		apiDto = &dto.ManualWithdrawsGrantPass{}
		apiRequestUrl = "manualWithdraws/grant/pass"
	case "/manualWithdraws/grantReject":
		apiDto = &dto.ManualWithdrawsGrantReject{}
		apiRequestUrl = "manualWithdraws/grant/reject"
	case "/providerWithdraws":
		apiDto = &dto.CreateProviderWithdraws{}
		apiRequestUrl = "providerWithdraws"
	case "/providerWithdraws/auditPass":
		apiDto = &dto.ProviderWithdrawsAuditPass{}
		apiRequestUrl = "providerWithdraws/audit/pass"
	case "/providerWithdraws/auditReject":
		apiDto = &dto.ProviderWithdrawsAuditReject{}
		apiRequestUrl = "providerWithdraws/audit/reject"
	case "/providerWithdraws/grantPass":
		apiDto = &dto.ProviderWithdrawsGrantPass{}
		apiRequestUrl = "providerWithdraws/grant/pass"
	case "/providerWithdraws/grantReject":
		apiDto = &dto.ProviderWithdrawsGrantReject{}
		apiRequestUrl = "providerWithdraws/grant/reject"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Octopus InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}