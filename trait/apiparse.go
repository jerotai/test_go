package trait

import (
	"github.com/gin-gonic/gin"
	"strings"
	"fmt"
	"io/ioutil"
	"reflect"
	"encoding/json"
	"net/http"
	"errors"
	"Stingray/core/model/redis/domain"
	"Stingray/helper"
	"Stingray/core/model/redis/rsakey"
	"mime/multipart"
	"bytes"
	"io"
)

/**
 * Station Type
 */
const (
	Site_Code  = "site_code"
	Hall_Code  = "hall_code"
)

type ApiResponse struct {
	Code string `json:"code"`
	Result interface{} `json:"result"`
	Message string `json:"message"`
}

type JSONInt struct {
	Value int
	Valid bool
	Set   bool
}

/**
 * 使用 *int 表示當前端未送參數時 將過濾 不加入傳送參數 (int 空值時 預設為0)
 */
func (i *JSONInt) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true
	
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	
	// The key isn't set to null
	var temp int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

type JSONString struct {
	Value string
	Valid bool
	Set   bool
}

/**
 * 使用 *string 表示當前端未送參數時 將過濾 不加入傳送參數 (int 空值時 預設為0)
 */
func (i *JSONString) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true
	
	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}
	
	// The key isn't set to null
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func GetToken(g *gin.Context) string {
	return g.Request.Header.Get("Api-Token")
}

func GetUserIp(g *gin.Context) string {
	client_ip := strings.Split(g.Request.RemoteAddr, ":")
	return client_ip[0]
}

/**
 * 取得前台user 呼叫 domian
 */
func GetUserDomain(g *http.Request, isPrefix bool) (string) {
	if g.Header.Get("Origin") != "" {
		domain_url := strings.Split(g.Header.Get("Origin"), "://")
		domainUrlPort := strings.Split(domain_url[1], ":")[0]
		domainUrl := strings.Split(domainUrlPort, "/")[0]
		
		helper.HelperLog.TraceLog(fmt.Sprintf("[parse-GetUserDomain] domain_url: %s", domain_url))
		helper.HelperLog.TraceLog(fmt.Sprintf("[parse-GetUserDomain] domainUrl: %s", domainUrl))
		
		if isPrefix == true {
			return domain_url[0] + "://" + domainUrl
		} else {
			return domainUrl
		}
	}
	
	if g.Header.Get("Referer") != "" {
		referer := strings.Split(g.Header.Get("Referer"), "://")
		var domainUrlPort = ""
		if len(referer) > 1 {
			domainUrlPort = strings.Split(referer[1], ":")[0]
			
		} else {
			domainUrlPort = strings.Split(g.Header.Get("Referer"), ":")[0]
		}
		
		domainUrl := strings.Split(domainUrlPort, "/")[0]
		
		helper.HelperLog.TraceLog(fmt.Sprintf("[parse-GetUserDomain] referer: %s", referer))
		helper.HelperLog.TraceLog(fmt.Sprintf("[parse-GetUserDomain] refererUrl: %s", domainUrl))
		
		if isPrefix == true {
			return referer[0] + "://" + domainUrl
		} else {
			return domainUrl
		}
	}
	
	return ""
}

/**
 * 取得 API GET 參數
 */
func DataParseByGet(g *gin.Context, apiDto interface{}) (apiMap map[string]interface{}, err error) {
	datas := make(map[string]interface{})
	apiMap = make(map[string]interface{})
	
	defer func(){
		if parseErr := recover(); parseErr != nil{
			err = errors.New(fmt.Sprintf(" [parse-DataParseByGet] : %s", parseErr))
			return
		}
	}()
	
	g.ShouldBindQuery(apiDto)
	
	for _, param := range g.Params {
		datas[param.Key] = param.Value
	}
	
	tmpJson, err := json.Marshal(datas)
	
	if err != nil {
		//後續處理
		err = errors.New(fmt.Sprintf(" [parse-DataParseByGet] : %s", err.Error()))
		return
	}
	
	json.Unmarshal(tmpJson, apiDto)
	
	if apiDto == nil {
		err = errors.New("[parse-DataParseByGet] apiDto is nil")
		return
	}
	
	if t := reflect.TypeOf(apiDto); t.Kind() == reflect.Ptr {
		
		v := reflect.ValueOf(apiDto)
		
		for i := 0; i < t.Elem().NumField(); i++ {
			fv := t.Elem().Field(i)
			tagMap := strings.Split(string(fv.Tag), ":")
			tag := fv.Tag.Get(tagMap[0]) //get struct tag type
			
			switch v.Elem().Field(i).Type().String() {
			case "int", "int32", "int64":
				apiMap[tag] = v.Elem().Field(i).Int()
			case "*int","*string":
				if v.Elem().Field(i).IsNil() != true {
					apiMap[tag] = v.Elem().Field(i).Elem()
				}
			case "float", "float32", "float64":
				apiMap[tag] = v.Elem().Field(i).Float()
			case "string":
				if(v.Elem().Field(i).String() != ""){
					apiMap[tag] = v.Elem().Field(i).String()
				}
			case "FileHeader":
				apiMap[tag] = v.Elem().Field(i).Bytes()
			default:
				apiMap[tag] = v.Elem().Field(i)
			}
		}
	}
	
	helper.HelperLog.TraceLog(fmt.Sprintf("[parse-DataParseByGet] apiMap: %s", apiMap))
	return
}

/**
 * 取得 API MultiPartDataParse 參數
 */
func MultiPartDataParse(g *gin.Context, apiDto interface{}, stationType string, stationCode string) (*bytes.Buffer, string) {
	var byteBuf bytes.Buffer
	mpWriter := multipart.NewWriter(&byteBuf)
	
	if t := reflect.TypeOf(apiDto); t.Kind() == reflect.Ptr {
		v := reflect.ValueOf(apiDto)
		
		for i := 0; i < t.Elem().NumField(); i++ {
			fv := t.Elem().Field(i)
			tagMap := strings.Split(string(fv.Tag), ":")
			tag := fv.Tag.Get(tagMap[0]) //get struct tag type
			switch v.Elem().Field(i).Type().String() {
			
			case "http.File":
				image, header, err := g.Request.FormFile(tag)
				if err == nil {
					if header != nil {
						fw, err := mpWriter.CreateFormFile(tag, header.Filename)
						
						if err != nil {
							helper.HelperLog.ErrorLog("[trait MultiPartDataParse mpWriter.CreateFormFile] " + err.Error())
						}
						
						if _, err = io.Copy(fw, image); err != nil {
							helper.HelperLog.ErrorLog("[trait MultiPartDataParse io.Copy ] " + err.Error())
						}
					}
				}
			case "int", "int32", "int64", "float", "float32", "float64":
				// Add the int fields
				if g.PostForm(tag) != "" {
					fw, _ := mpWriter.CreateFormField(tag)
					fw.Write([]byte(g.PostForm(tag)))
				}
			default:
				// Add the other fields
				fw, _ := mpWriter.CreateFormField(tag)
				fw.Write([]byte(g.PostForm(tag)))
			}
		}
	}
	
	mpWriter.WriteField(stationType, stationCode)

	mpWriter.Close()
	
	return &byteBuf, mpWriter.FormDataContentType()
}


/**
 * 取得 API POST PUT DELETE 參數
 */
func DataParse(g *gin.Context, apiDto interface{}) (apiMap map[string]interface{}, err error) {
	var inputData = make(map[string]interface{})
	apiMap = make(map[string]interface{})
	
	defer func(){
		if parseErr := recover(); parseErr != nil{
			err = errors.New(fmt.Sprintf("[parse-DataParse] : %s", parseErr))
			return
		}
	}()
	
	if apiDto == nil {
		err = errors.New("[parse-DataParse] apiDto is nil")
		return
	}
	
	body, err := ioutil.ReadAll(g.Request.Body)
	helper.HelperLog.TraceLog("[parse-DataParse] body: " + string(body))
	
	if err != nil {
		//後續處理
		err = errors.New("[parse-DataParse] ioutil.ReadAll error " + err.Error())
		return
	}
	
	json.Unmarshal(body, &inputData)
	
	apiMap = conformDataParse(apiDto, inputData)
	
	helper.HelperLog.TraceLog(fmt.Sprintf("[parse-DataParse] Unmarshal apiDto: %s", apiMap))
	return
}

/**
 * 取得 rsa加密的post,put,delete 參數
 */
func DataParseByRsa(g *gin.Context, apiDto interface{}) (apiMap map[string]interface{}, err error) {
	var inputData = make(map[string]interface{})
	apiMap = make(map[string]interface{})
	
	defer func(){
		if parseErr := recover(); parseErr != nil{
			err = errors.New(fmt.Sprintf(" [parse-DataParseByRsa] : %s", parseErr))
			return
		}
	}()
	
	body, err := ioutil.ReadAll(g.Request.Body)
	helper.HelperLog.TraceLog("[parse-DataParseByRsa] Input Body: " + string(body))
	
	if err != nil {
		//後續處理
		err = errors.New(fmt.Sprintf(" [parse-DataParseByRsa] : Body ", err.Error()))
		return
	}
	
	//取得指定 token 的 rsa private key
	token := g.Request.Header.Get("Api-Token")
	rsaRedisService := rsakey.New()
	rsaKey := rsaRedisService.GetTokenRsaPrivateKey(token)
	
	if rsaKey == "" {
		return apiMap, errors.New("[parse-DataParseByRsa] : Rsa Key Is Nil")
	}
	
	//input rsa encode
	rsadecode, err := helper.RsaDecryptByKey(rsaKey, body)
	
	helper.HelperLog.TraceLog(fmt.Sprintf("[parse-DataParseByRsa] RsaDecryptByKey body: %s", rsadecode))

	if err != nil {
		//todo 錯誤處理
		return apiMap, errors.New("[parse-DataParseByRsa] Error: " + err.Error())
	}
	
	json.Unmarshal(rsadecode, &inputData)
	
	apiMap = conformDataParse(apiDto, inputData)

	helper.HelperLog.TraceLog(fmt.Sprintf("[parse-DataParseByRsa] Unmarshal apiMap: %s", apiMap))
	return
}

/**
 * 整合 api input data by apiDto
 */
func conformDataParse(apiDto interface{}, inputData map[string]interface{}) (map[string]interface{}) {
	var apiMap = make(map[string]interface{})
	
	if t := reflect.TypeOf(apiDto); t.Kind() == reflect.Ptr {
		
		v := reflect.ValueOf(apiDto)
		
		for i := 0; i < t.Elem().NumField(); i++ {
			fv := t.Elem().Field(i)
			tagMap := strings.Split(string(fv.Tag), ":")
			tag := fv.Tag.Get(tagMap[0]) //get struct tag type
			
			//如果有特別的型別需要確認 可以加上特別處理方式
			switch v.Elem().Field(i).Type().String() {
			case "int", "int32", "int64":
				if(inputData[tag] != nil) {
					apiMap[tag] = int(inputData[tag].(float64))
				}
			case "float", "float32", "float64":
				if(inputData[tag] != nil) {
					apiMap[tag] = fmt.Sprintf("%f", inputData[tag])
				}
			default:
				if(inputData[tag] != nil) {
					apiMap[tag] = inputData[tag]
				}
			}
		}
	}
	
	return apiMap
}

/**
 * 確認站台 Code
 * url := StationCodeParse(g.Request)
 */
func StationCodeParse(g *http.Request) string {
	//start domain list redis service
	domain := domain.DomainsService()
	
	domainUrl := GetUserDomain(g, false)
	fmt.Println("domainUrl", domainUrl)
	return domain.GetStationCodeByDomain(domainUrl)
}
