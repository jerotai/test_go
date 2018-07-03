package domain

import "Stingray/helper"

/**
 * 取得 指定 domain url site code
 */
func (d *Domain) GetStationCodeByDomain(url string) string {
	code := d.data_connect.HGet("domain_list", url)
	helper.HelperLog.TraceLog("[domain GetStationCodeByDomain] url : " + url)
	helper.HelperLog.TraceLog("[domain GetStationCodeByDomain] code : " + code.Val())
	return code.Val()
}