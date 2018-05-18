package domain

/**
 * 取得 指定 domain url site code
 */
func (d *Domain) GetStationCodeByDomain(url string) string {
	//d.data_connect.Get(url)
	//get redis connect
	code := d.data_connect.HGet("domain_list", url)
	//check domain code
	
	return code.Val()
}