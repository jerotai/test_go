package apiparse

import (
	"encoding/json"
	"net/http"
	"strings"
	"reflect"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"Stingray/core/model/redis/domain"
	"Stingray/core/dto"
	"fmt"
	"Stingray/helper"
	"Stingray/core/model/redis/rsakey"
)

/**
 * 回傳格式彙整
 */
func ApiResponse(req string) dto.ApiStatus {
	statusReq := dto.ApiStatus{}
	json.Unmarshal([]byte(req), &statusReq)
	return statusReq
}

/**
 * 取得 API GET 參數
 */
func GetDataParse(g *gin.Context, apiDto interface{}) error {
	datas := make(map[string]interface{})
	/*decoder := schema.NewDecoder()
	err := decoder.Decode(apiDto, g.Request.URL.Query())
	if err != nil {
		//後續處理
		fmt.Println(err)
	}
	fmt.Println("url query decode to dto : ", apiDto)*/
	g.ShouldBindQuery(apiDto)
	for _, param := range g.Params {
		datas[param.Key] = param.Value
	}
	
	tmpJson, err := json.Marshal(datas)
	fmt.Println("Marshal : ", string(tmpJson))
	if err != nil {
		//後續處理
	}
	json.Unmarshal(tmpJson, apiDto)
	fmt.Println("Unmarshal : ", apiDto)
	return err
}

/**
 * 取得 API POST 參數
 */
func PostDataParse(g *gin.Context, apiDto interface{}) error {
	body, err := ioutil.ReadAll(g.Request.Body)
	if err != nil {
		//後續處理
	}
	//rsadecode, _ := helper.RsaDecrypt(body)
	json.Unmarshal(body, apiDto)
	fmt.Println("apiDto", apiDto)
	return err
}

/**
 * 取得 rsa加密的post 參數
 */
func RsaDataParse(g *gin.Context, apiDto interface{}) error {
	body, err := ioutil.ReadAll(g.Request.Body)
	fmt.Println("def", string(body))
	if err != nil {
		//後續處理
	}
	
	//取得指定 token 的 rsa private key
	token := g.Request.Header.Get("Api-Token")
	rsaRedisService := rsakey.RsaKeyService()
	rsaRedisService.Init()
	rsaKey := rsaRedisService.GetTokenRsaPrivateKey(token)
	
	//input rsa encode
	rsadecode, _ := helper.RsaDecryptByKey(rsaKey, body)
	//rsadecode, _ := helper.RsaDecrypt(body)
	
	fmt.Println("de", string(rsadecode))
	json.Unmarshal(rsadecode, apiDto)
	
	return err
}

/**
 * 將 2層dto Struct 轉換為 map interface
 */
func TowLayerDtoToMap(apiDto interface{}) (apiMap map[string]interface{}, err error) {
	apiMap = make(map[string]interface{})
	
	if apiDto == nil {
		err = errors.New("apiDto is nil")
	}

	if t := reflect.TypeOf(apiDto); t.Kind() == reflect.Ptr {
		
		v := reflect.ValueOf(apiDto)
		
		for i := 0; i < t.Elem().NumField(); i++ {
			fv := t.Elem().Field(i)
			tag := fv.Tag.Get("json")
			switch v.Elem().Field(i).Type().String() {
			case "int":
				apiMap[tag] = v.Elem().Field(i).Int()
			case "string":
				apiMap[tag] = v.Elem().Field(i).String()
			default:
				apiMap[tag] = v.Elem().Field(i)
			}
		}
	}
	
	err = errors.New("illegal input")
	return
}

/**
 * 將 dto Struct 轉換為 map interface
 */
func DtoToMap(apiDto interface{}) (apiMap map[string]interface{}, err error) {
	apiMap = make(map[string]interface{})
	
	if apiDto == nil {
		err = errors.New("apiDto is nil")
	}

	if t := reflect.TypeOf(apiDto); t.Kind() == reflect.Struct {
		
		v := reflect.ValueOf(apiDto)
		
		for i := 0; i < t.NumField(); i++ {
			fv := t.Field(i)
			tag := fv.Tag.Get("json")
			
			switch v.Field(i).Type().String() {
			case "int":
				apiMap[tag] = v.Field(i).Int()
			case "string":
				apiMap[tag] = v.Field(i).String()
			default:
				apiMap[tag] = v.Field(i)
			}
		}
	}
	
	err = errors.New("illegal input")
	return
}

/**
 * 確認站台 Code
 * url := StationCodeParse(g.Request)
 */
func StationCodeParse(g *http.Request) string {
	
	if g.Header.Get("origin") != "" {
		domain_url := strings.Split(g.Header.Get("origin"), "://")
		domain := domain.DomainsService()
		domain.Init()
		domainUrl := strings.Split(domain_url[1], ":")[0]
		return domain.GetStationCodeByDomain(domainUrl)
	}
	
	return "cqcp"
}

