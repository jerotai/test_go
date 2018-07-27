package whitebaitdto

/**
 * Api Url : guest/login (POST)
 */
type GuestLogin struct {
	Plat_Id int `json:"plat_id"`
}

/**
 * Api Url : guest/myself (GET)取得訪客設定
 */
type GuestMyself struct {
}

/**
 * Api Url : guest/myself (PUT) 更新訪客設定
 */
type UpdateGuestMyself struct {
	Balance int `json:"balance"`
	Visitor int `json:"visitor"`
	Order int `json:"order"`
}

/**
 * Api Url : guest/hall (GET) 取得下層廳主訪客設定
 */
type GuestHall struct {
	Id int `json:"id"`
}

/**
 * Api Url : guest/hall (PUT) 更新下層廳主訪客設定
 */
type UpdateGuestHall struct {
	Id int `json:"id"`
	Balance int `json:"balance"`
	Visitor int `json:"visitor"`
	Order int `json:"order"`
}

/**
 * Api Url : guest/site (GET) 取得站台訪客設定
 */
type GuestSite struct {
	Site_Code string `json:"site_code"`
}

/**
 * Api Url : guest/site (PUT) 更新站台訪客設定
 */
type UpdateGuestSite struct {
	Site_Code string `json:"site_code"`
	Balance int `json:"balance"`
	Visitor int `json:"visitor"`
	Order int `json:"order"`
}