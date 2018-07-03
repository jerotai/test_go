package jellyfishdto

/**
 * Api Url : agentSystem/siteTotal (GET) 站台下層體系數量清單
 */
type AgentSystemSiteTotal struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : agentSystem/shareholderTotal (GET) 站台股東下層體系數量清單
 */
type AgentSystemShareholderTotal struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Shareholder_Name string `json:"shareholder_name"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : agentSystem/generalAgentTotal (GET) 站台總代理下層體系數量清單
 */
type AgentSystemGeneralAgentTotal struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Shareholder_Name string `json:"shareholder_name"`
	General_Agent_Name string `json:"general_agent_name"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : agentSystem/agentTotal (GET) 站台代理下層體系數量清單
 */
type AgentSystemAgentTotal struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	General_Agent_Name string `json:"general_agent_name"`
	Agent_Name string `json:"agent_name"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : agentSystem/memberTotal (GET) 站台會員清單
 */
type AgentSystemMemberTotal struct {
	Hall_Code string `json:"hall_code"`
	Site_Code string `json:"site_code"`
	Shareholder_Name string `json:"shareholder_name"`
	General_Agent_Name string `json:"general_agent_name"`
	Agent_Name string `json:"agent_name"`
	Member_Name string `json:"member_name"`
	Page string `json:"page"`
	Count string `json:"count"`
}