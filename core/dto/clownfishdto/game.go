package clownfishdto

/**
 * Api Url : game/provider/list (GET) 取得遊戲商清單
 */
 type GameProviderList struct {
}

/**
 * Api Url : game/provider/kind (GET) 取得遊戲商的遊戲類別清單
 */
 type GameProviderKind struct {
	Id string `json:"id"`
}