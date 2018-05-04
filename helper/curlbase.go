package helper

import (
	//"reflect"
	//"routes/core/dto"
	"errors"
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type CurlBase struct {
	Url  string
}

// CombineParamUri
//
// 所有參數轉為uri格式字串,參數依照key值排序
// ex. hash=asd&param1=a&param2=b
func combineParamUri(param map[string]interface{}) (string, error) {
	if param == nil {
		return "", nil
	}

	u := url.URL{}

	//add paramter
	query := u.Query()
	for k, v := range param {
		query.Set(k, fmt.Sprint(v))
	}

	return query.Encode(), nil
}

// warnErrorData
//
// 當API回傳的code != 0時，代表有異常發生
// 外部收到error後再依照error code做對應處理
func (inst CurlBase) warnErrorData(resObj interface{}) error {

	var message string
	/*if t := reflect.ValueOf(resObj); t.Kind() == reflect.Ptr {
		// 外部傳入的interface是 ** 所以要Elem兩次
		f := t.Elem().Elem().FieldByName("Status")
		s := f.Interface().(dto.Status)
		if s.Code == "0" {
			return nil
		}
		message = s.Message
	}*/
	return errors.New(message)
}

func (inst CurlBase) Get(path string, params map[string]interface{}, resObj interface{}) error {
	client := &http.Client{}
	paramUri, err := combineParamUri(params)

	if err != nil {
		return err
	}
	fmt.Println(paramUri)

	apiUrl := inst.Url + path + "?" + paramUri

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)

	if err != nil {
		return errors.New(
			fmt.Sprintf(
				"Msg:%s api:%s err:%s",
				"NewRequest Failed",
				apiUrl,
				err.Error(),
			),
		)
	}

	//req.Header.Add("Authorization", inst.Auth)

	res, err := client.Do(req)
	if err != nil {
		return errors.New(
			fmt.Sprintf(
				"Msg:%s api:%s err:%s",
				"Send Request Failed",
				apiUrl,
				err.Error(),
			),
		)
	}

	//close connection
	defer res.Body.Close()

	//read all response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New(
			fmt.Sprintf(
				"Msg:%s api:%s err:%s",
				"Read From body Failed",
				apiUrl,
				err.Error(),
			),
		)
	}

	//Try Unmarshal Json
	err = json.Unmarshal(body, resObj)
	if err != nil {
		return errors.New(
			fmt.Sprintf(
				"Msg:%s API:%s Res:%s err:%s",
				"Unmarshal Json Failed",
				apiUrl,
				string(body),
				err.Error(),
			),
		)
	}
	return inst.warnErrorData(resObj)
}