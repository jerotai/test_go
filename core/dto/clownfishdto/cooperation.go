package clownfishdto

/**
 * Api Url : cooperation/fortune (GET) 取得好運來連結（未登入）
 */
type CooperationFortune struct {
	Code string `json:"code"`
}

/**
 * Api Url : cooperation/fortune/user (GET) 取得好運來連結
 */
type CooperationFortuneUser struct {
	Code string `json:"code"`
}