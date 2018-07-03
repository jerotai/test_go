package octopus

import (
	dto "Stingray/core/dto/octopusdto"
	
	"Stingray/helper"
)

/**
 * init Api Put Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Octopus) InitPutApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/trans/unlock":
		apiDto = &dto.TransUnlock{}
		apiRequestUrl = "trans/unlock"
	case "/trans/selfunlock":
		apiDto = &dto.TransSelfUnlock{}
		apiRequestUrl = "trans/selfunlock"
	case "/manualDeposits/unlock":
		apiDto = &dto.ManualDepositsUnlock{}
		apiRequestUrl = "manualDeposits/unlock"
	case "/fourthDeposits/unlock":
		apiDto = &dto.FourthDepositsUnlock{}
		apiRequestUrl = "fourthDeposits/unlock"
	case "/bankDeposits/unlock":
		apiDto = &dto.BankDepositsUnlock{}
		apiRequestUrl = "bankDeposits/unlock"
	case "/providerDeposits/unlock":
		apiDto = &dto.ProviderDepositsUnlock{}
		apiRequestUrl = "providerDeposits/unlock"
	case "/manualWithdraws/unlock":
		apiDto = &dto.ManualWithdrawsUnlock{}
		apiRequestUrl = "manualWithdraws/unlock"
	case "/providerWithdraws/unlock":
		apiDto = &dto.ProviderWithdrawsUnlock{}
		apiRequestUrl = "providerWithdraws/unlock"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Octopus InitPutApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}