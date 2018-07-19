package electricity


import (
	dto "Stingray/core/dto/electricitydto"
	"Stingray/helper"
)

/**
 * init Api Get Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (i *Electricity) InitGetApiConfig(apiUrl string) (string, interface{}){
	apiRequestUrl := ""
	var apiDto interface {}
	switch apiUrl {
	case "/reward_config":
		apiDto = &dto.RewardConfig{}
		apiRequestUrl = "reward_config"
	case "/reward/info":
		apiDto = &dto.RewardInfo{}
		apiRequestUrl = "reward/info"
	case "/reward/record":
		apiDto = &dto.RewardRecord{}
		apiRequestUrl = "reward/record"
	case "/reward/dropdown":
		apiDto = &dto.RewardDropdown{}
		apiRequestUrl = "reward/dropdown"
	case "/reward/list":
		apiDto = &dto.RewardList{}
		apiRequestUrl = "reward"
	case "/reward/dropdownList":
		apiDto = &dto.RewardDropdownlist{}
		apiRequestUrl = "reward/dropdownlist"
	case "/reward/reject":
		apiDto = &dto.RewardReject{}
		apiRequestUrl = "reward/reject"
	case "/reward/pass":
		apiDto = &dto.RewardPass{}
		apiRequestUrl = "reward/pass"
	case "/reward/report":
		apiDto = &dto.RewardReport{}
		apiRequestUrl = "reward/report"
	case "/reward/reportInvitee":
		apiDto = &dto.RewardReportInvitee{}
		apiRequestUrl = "reward/report/invitee"
	default:
		//todo
		helper.HelperLog.ErrorLog("[Electricity InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
	}
	
	return apiRequestUrl, apiDto
}
