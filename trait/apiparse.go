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

func GetToken(g *gin.Context) string {
	return g.Request.Header.Get("Api-Token")
}

func GetUserIp(g *gin.Context) string {
	client_ip := strings.Split(g.Request.RemoteAddr, ":")
	return client_ip[0]
}

/**
 * 取得 API GET 參數
 */
func DataParseByGet(g *gin.Context, apiDto interface{}) error {
	datas := make(map[string]interface{})
	/*
	decoder := schema.NewDecoder()
	err := decoder.Decode(apiDto, g.Request.URL.Query())
	if err != nil {
		//後續處理
		fmt.Println(err)
	}
	*/
	g.ShouldBindQuery(apiDto)
	
	for _, param := range g.Params {
		datas[param.Key] = param.Value
	}
	
	tmpJson, err := json.Marshal(datas)
	
	if err != nil {
		//後續處理
		helper.HelperLog.ErrorLog("[parse-DataParseByGet] " + g.Request.URL.Path + ": " + err.Error())
	}
	json.Unmarshal(tmpJson, apiDto)
	
	return err
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
			case "int":
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
func DataParse(g *gin.Context, apiDto interface{}) error {
	body, err := ioutil.ReadAll(g.Request.Body)
	helper.HelperLog.TraceLog("[parse-DataParse] body: " + string(body))
	if err != nil {
		//後續處理
		helper.HelperLog.ErrorLog(err.Error())
	}
	
	json.Unmarshal(body, apiDto)
	helper.HelperLog.TraceLog(fmt.Sprintf("[parse-DataParse] Unmarshal apiDto: %s", apiDto))
	return err
}

/**
 * 取得 rsa加密的post 參數
 */
func DataParseByRsa(g *gin.Context, apiDto interface{}) error {
	body, err := ioutil.ReadAll(g.Request.Body)
	helper.HelperLog.TraceLog("[parse-DataParseByRsa] body: " + string(body))
	if err != nil {
		//後續處理
		helper.HelperLog.ErrorLog("[parse-DataParseByRsa] Body" + g.Request.URL.Path + ": " + err.Error())
	}
	
	//取得指定 token 的 rsa private key
	token := g.Request.Header.Get("Api-Token")
	rsaRedisService := rsakey.New()
	rsaKey := rsaRedisService.GetTokenRsaPrivateKey(token)
	
	if rsaKey == "" {
		return errors.New("Rsa Key Is Nil")
	}
	
	//input rsa encode
	rsadecode, err := helper.RsaDecryptByKey(rsaKey, body)
	helper.HelperLog.TraceLog(fmt.Sprintf("[parse-DataParseByRsa] RsaDecryptByKey body: %s", string(body)))

	if err != nil {
		//todo 錯誤處理
		helper.HelperLog.ErrorLog("[parse-DataParseByRsa] Error: " + g.Request.URL.Path + ": " + err.Error())
		return errors.New(err.Error())
	}
	
	json.Unmarshal(rsadecode, apiDto)
	helper.HelperLog.TraceLog(fmt.Sprintf("[parse-DataParseByRsa] Unmarshal apiDto: %s", apiDto))
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
			tagMap := strings.Split(string(fv.Tag), ":")
			tag := fv.Tag.Get(tagMap[0]) //get struct tag type
			
			switch v.Elem().Field(i).Type().String() {
			case "int":
				apiMap[tag] = v.Elem().Field(i).Int()
			case "int32":
				apiMap[tag] = v.Elem().Field(i).Int()
			case "int64":
				apiMap[tag] = v.Elem().Field(i).Int()
			case "float":
				apiMap[tag] = v.Elem().Field(i).Float()
			case "float32":
				apiMap[tag] = v.Elem().Field(i).Float()
			case "float64":
				apiMap[tag] = v.Elem().Field(i).Float()
			case "string":
				apiMap[tag] = v.Elem().Field(i).String()
			case "FileHeader":
				apiMap[tag] = v.Elem().Field(i).Bytes()
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
			tagMap := strings.Split(string(fv.Tag), ":")
			tag := fv.Tag.Get(tagMap[0]) //get struct tag type
			
			switch v.Field(i).Type().String() {
			case "int":
				apiMap[tag] = v.Field(i).Int()
			case "int32":
				apiMap[tag] = v.Field(i).Int()
			case "int64":
				apiMap[tag] = v.Field(i).Int()
			case "float":
				apiMap[tag] = v.Field(i).Float()
			case "float32":
				apiMap[tag] = v.Field(i).Float()
			case "float64":
				apiMap[tag] = v.Field(i).Float()
			case "string":
				apiMap[tag] = v.Field(i).String()
			case "FileHeader":
				apiMap[tag] = v.Elem().Field(i).String()
			default:
				apiMap[tag] = v.Field(i)
			}
		}
	}
	
	return
}

/**
 * 確認站台 Code
 * url := StationCodeParse(g.Request)
 */
func StationCodeParse(g *http.Request) string {
	//start domain list redis service
	domain := domain.DomainsService()
	
	if g.Header.Get("origin") != "" {
		domain_url := strings.Split(g.Header.Get("Origin"), "://")
		helper.HelperLog.TraceLog(fmt.Sprintf("[parse-StationCodeParse] domain_url: %s", domain_url))
		domainUrl := strings.Split(domain_url[1], ":")[0]
		helper.HelperLog.TraceLog(fmt.Sprintf("[parse-StationCodeParse] domainUrl: %s", domainUrl))
		return domain.GetStationCodeByDomain(domainUrl)
	}
	
	if g.Header.Get("Referer") != "" {
		referer := strings.Split(g.Header.Get("Referer"), "://")
		refererUrl := ""
		if len(referer) > 1 {
			refererUrl = strings.Split(referer[1], ":")[0]
			
		} else {
			refererUrl = strings.Split(g.Header.Get("Referer"), ":")[0]
		}
		helper.HelperLog.TraceLog(fmt.Sprintf("[parse-StationCodeParse] refererUrl: %s", refererUrl))
		return domain.GetStationCodeByDomain(refererUrl)
	}
	
	return ""
}
