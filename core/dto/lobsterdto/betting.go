package lobsterdto

/**
 * Api Url : betting/subordinate (GET) 下線投注報表-清單
 */
type BettingSubordinate struct {
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Sort_By string `json:"sort_by"`
	Order string `json:"order"`
}

/**
 * Api Url : betting/subordinate/summary (GET) 下線投注報表-總計
 */
type BettingSubordinateSummary struct {
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
}

/**
 * Api Url : betting/game (GET) 遊戲投注報表
 */
type BettingGame struct {
	Site_Code string `json:"site_code"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Page string `json:"page"`
	Count string `json:"count"`
	Order string `json:"order"`
	Orderby string `json:"orderby"`
}