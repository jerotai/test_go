package jellyfishdto

/**
 * Api Url : agentSub (GET)
 */
type AgentSubList struct {
	Account string `json:"account"`
	Status string `json:"status"`
	Name string `json:"name"`
	Page string `json:"page"`
	Count string `json:"count"`
}

/**
 * Api Url : agentSub/data (GET)
 */
type AgentSubData struct {
	Id string `json:"id"`
}

/**
 * Api Url : agentSub (POST)
 */
type CreateAgentSub struct {
	Agent_id int `json:"agent_id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Authority int `json:"authority"`
}

/**
 * Api Url : agentSub (PUT)
 */
type UpdateAgentSub struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Authority int `json:"authority"`
	Status int `json:"status"`
}

/**
 * Api Url : agentSub/password 更新代理子帳號密碼
 */
type UpdateAgentSubPassword struct {
	Id int `json:"id"`
	Password string `json:"password"`
}