package electricitydto

/**
 * Api Url : reward_config (GET) 取得站點推廣退傭設定
 */
type RewardConfig struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : reward_config (PUT) 更新站點推廣退傭設定
 */
type UpdateRewardConfig struct {
	Site_Code string `json:"site_code"`
	Reward_Enable int `json:"reward_enable"`
	Reward_Permille int `json:"reward_permille"`
	Upper_Limit float32 `json:"upper_limit"`
	Lower_Limit float32 `json:"lower_limit"`
	Settle_Type int `json:"settle_type"`
	Settle_Day int `json:"settle_day"`
	Settle_Hour int `json:"settle_hour"`
}

/**
 * Api Url : reward/info (GET) 取得會員推廣資訊
 */
type RewardInfo struct {
}

/**
 * Api Url : reward/record (GET) 取得會員佣金紀錄
 */
type RewardRecord struct {
	Issue string `json:"issue"`
}

/**
 * Api Url : reward/dropdown (GET) 取得站點最新5筆期數清單（下拉選單）
 */
type RewardDropdown struct {
}

/**
 * Api Url : reward (GET) 取得站點推廣退傭清單
 */
type RewardList struct {
	Site_Code string `json:"site_code"`
	Status string `json:"status"`
	Issue string `json:"issue"`
	Referrer string `json:"referrer"`
	User string `json:"user"`
	Level string `json:"level"`
	Amount string `json:"amount"`
	Count string `json:"count"`
	Page string `json:"page"`
}

/**
 * Api Url : reward/dropdownlist (GET)取得站點推廣期數清單（下拉選單）
 */
type RewardDropdownlist struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : reward/reject (GET) 推廣退傭駁回
 */
type RewardReject struct {
	Items string `json:"items"`
	Remark string `json:"remark"`
}

/**
 * Api Url : reward/pass (GET) 推廣退傭派發通過
 */
type RewardPass struct {
	Items string `json:"items"`
	Dama_Ratio string `json:"dama_ratio"`
	Remark string `json:"remark"`
}

/**
 * Api Url : reward/report (GET) 取得推廣會員統計清單（上層）
 */
type RewardReport struct {
	site_code string `json:"site_code"`
	Referrer string `json:"referrer"`
	Count string `json:"count"`
	Page string `json:"page"`
	Sort_By string `json:"sort_by"`
	Order string `json:"order"`
}

/**
 * Api Url : reward/report/invitee (GET) 取得被推廣會員統計清單（下層）
 */
type RewardReportInvitee struct {
	site_code string `json:"site_code"`
	Referrer_Id string `json:"referrer_id"`
	Count string `json:"count"`
	Page string `json:"page"`
	Sort_By string `json:"sort_by"`
	Order string `json:"order"`
}