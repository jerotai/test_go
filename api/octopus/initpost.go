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
	case "/providerWithdraws":
		apiDto = &dto.CreateProviderWithdraws{}
		apiRequestUrl = "providerWithdraws"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Octopus InitPostApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}