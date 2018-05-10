package apiparse

import (
	"encoding/json"
	"routes/core/dto"
	"net/http"
	"routes/core/model/redis/domain"
	"strings"
	"reflect"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"routes/helper"
)

/**
 * 回傳格式彙整
 */
func ApiResponse(req interface{}) interface{} {
	statusReq := dto.ApiStatus{}
	b, _ := json.Marshal(req)
	
	//statusReq.Code = statusCodeParse("1")
	statusReq.Result = string(b[:])
	return statusReq
}


/**
 * 取得 API GET 參數
 */
func GetDataParse(g *gin.Context, apiDto interface{}) error {
	datas := make(map[string]interface{})
	
	for _, param := range g.Params {
		datas[param.Key] = param.Value
	}
	
	tmpJson, err := json.Marshal(datas)
	
	if err != nil {
		//後續處理
	}
	
	json.Unmarshal([]byte(tmpJson), apiDto)
	
	return err
}

/**
 * 取得 API POST 參數
 */
func PostDataParse(g *gin.Context, apiDto interface{}) error {
	err := g.ShouldBind(apiDto)
	
	if err != nil {
		//後續處理
	}
	
	return err
}

/**
 * 取得 rsa加密的post 參數
 */
func PostRsaDataParse(g *gin.Context, apiDto interface{}) error {
	body, err := ioutil.ReadAll(g.Request.Body)
	
	if err != nil {
		//後續處理
	}
	rsadecode, _ := helper.RsaDecrypt(body)
	json.Unmarshal([]byte(rsadecode), apiDto)
	
	return err
}

/**
 * 將 dto Struct 轉換為 map interface
 */
func DtoToMapInterface(apiDto interface{}) (apiMap map[string]interface{}, err error) {
	apiMap = make(map[string]interface{})
	
	if apiDto == nil {
		err = errors.New("apiDto is nil")
	}
	
	if t := reflect.TypeOf(apiDto); t.Kind() == reflect.Struct {
		
		v := reflect.ValueOf(apiDto)
		
		for i := 0; i < t.NumField(); i++ {
			fv := t.Field(i)
			tag := fv.Tag.Get("json")
			apiMap[tag] = v.Field(i)
		}
	}
	
	err = errors.New("illegal input")
	return
}

/**
 * 確認站台 Code
 * url := siteCodeParse(g.Request)
 */
func SiteCodeParse(g *http.Request) string {
	domain_url := strings.Split(g.Host, ":")
	
	domain := domain.DomainsService()
	domain.Init()

	return domain.GetSiteCodeByDomain(domain_url[0])
}

