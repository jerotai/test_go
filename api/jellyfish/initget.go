package jellyfish

import (
	dto "Stingray/core/dto/jellyfishdto"
	"Stingray/helper"
)

/**
 * init Api Get Config
 * i.ApiConf.dto = &dto.Menu{}
 * i.ApiConf.apiRequestUrl = "menu"
 *
 */
func (j *JellyFish) InitGetApiConfig(apiUrl string) (string, interface{}) {
	apiRequestUrl := ""
	var apiDto interface {}
	
	switch apiUrl {
	case "/role/list":
		apiDto = &dto.RoleList{}
		apiRequestUrl = "role"
	case "/hall/subList":
		apiDto = &dto.SubList{}
		apiRequestUrl = "hall/subList"
	case "/hall/dropdownlist":
		apiDto = &dto.HallDropdownList{}
		apiRequestUrl = "hall/dropdownlist"
	case "/shareholder/list":
		apiDto = &dto.ShareholderList{}
		apiRequestUrl = "shareholder"
	case "/shareholder/data":
		apiDto = &dto.ShareholderData{}
		apiRequestUrl = "shareholder/data"
	case "/shareholder/dropdownlist":
		apiDto = &dto.ShareholderDropdownList{}
		apiRequestUrl = "shareholder/dropdownlist"
	case "/shareholderSub/list":
		apiDto = &dto.ShareholderSubList{}
		apiRequestUrl = "shareholderSub"
	case "/shareholderSub/data":
		apiDto = &dto.ShareholderSubData{}
		apiRequestUrl = "shareholderSub/data"
	case "/hallSub/list":
		apiDto = &dto.HallSubList{}
		apiRequestUrl = "hallSub"
	case "/hallSub/data":
		apiDto = &dto.HallSubData{}
		apiRequestUrl = "hallSub/data"
	case "/authGroup/list":
		apiDto = &dto.AuthGroupList{}
		apiRequestUrl = "authGroup"
	case "/authGroup/detail":
		apiDto = &dto.AuthGroupDatail{}
		apiRequestUrl = "authGroup/detail"
	case "/authGroup/dropDown":
		apiDto = &dto.AuthGroupDropDown{}
		apiRequestUrl = "authGroup/dropDown"
	case "/generalAgent/list":
		apiDto = &dto.GeneralAgentList{}
		apiRequestUrl = "generalAgent"
	case "/generalAgent/data":
		apiDto = &dto.GeneralAgentData{}
		apiRequestUrl = "generalAgent/data"
	case "/generalAgent/dropdownlist":
		apiDto = &dto.GeneralAgentDropdownList{}
		apiRequestUrl = "generalAgent/dropdownlist"
	case "/generalAgentSub/list":
		apiDto = &dto.GeneralAgentSubList{}
		apiRequestUrl = "generalAgentSub"
	case "/generalAgentSub/data":
		apiDto = &dto.GeneralAgentSubData{}
		apiRequestUrl = "generalAgentSub/data"
	case "/agent/list":
		apiDto = &dto.AgentList{}
		apiRequestUrl = "agent"
	case "/agent/data":
		apiDto = &dto.AgentData{}
		apiRequestUrl = "agent/data"
	case "/agent/dropdownlist":
		apiDto = &dto.AgentDropdownList{}
		apiRequestUrl = "agent/dropdownlist"
	case "/agentSub/list":
		apiDto = &dto.AgentSubList{}
		apiRequestUrl = "agentSub"
	case "/agentSub/data":
		apiDto = &dto.AgentSubData{}
		apiRequestUrl = "agentSub/data"
	case "/agentSystem/siteTotal":
		apiDto = &dto.AgentSystemSiteTotal{}
		apiRequestUrl = "agentSystem/siteTotal"
	case "/agentSystem/shareholderTotal":
		apiDto = &dto.AgentSystemShareholderTotal{}
		apiRequestUrl = "agentSystem/shareholderTotal"
	case "/agentSystem/generalAgentTotal":
		apiDto = &dto.AgentSystemGeneralAgentTotal{}
		apiRequestUrl = "agentSystem/generalAgentTotal"
	case "/agentSystem/agentTotal":
		apiDto = &dto.AgentSystemAgentTotal{}
		apiRequestUrl = "agentSystem/agentTotal"
	case "/agentSystem/memberTotal":
		apiDto = &dto.AgentSystemMemberTotal{}
		apiRequestUrl = "agentSystem/memberTotal"
	default:
		//todo
		helper.HelperLog.ErrorLog("[JellyFish InitGetApiConfig] Url Not Merge : apiUrl: " + apiUrl)
		}
	
	return apiRequestUrl, apiDto
}
