package example

import (
	"testing"
	"net/http/httptest"
	"github.com/labstack/echo"
	"net/http"
	"routes/api/example"
	"strings"
	"routes/helper"
	"github.com/stretchr/testify/assert"
)

var (
	exampleJSOM = `{"id":"111","name":"Jon Snow"}`
	reExampleJSOM = `{"code":"1","result":"{\"id\":\"111\",\"name\":\"Jon Snow\"}"}`
)

/**
 * 測試 GET Info
 */
func TestGetInfo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/example/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/example/:id/:name")
	c.SetParamNames("id", "name")
	c.SetParamValues("111", "Jon Snow")
	
	ex := &example.Example{}
	if assert.NoError(t, ex.GetInfo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, reExampleJSOM, rec.Body.String())
	}
}

/**
 *測試 POST
 */
func TestPost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/example", strings.NewReader(exampleJSOM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	
	c := e.NewContext(req, rec)
	
	ex := &example.Example{}
	
	if assert.NoError(t, ex.Post(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, reExampleJSOM, rec.Body.String())
	}
}

/**
 * 測試 RSA POST
 */
func TestRsa(t *testing.T) {
	rsaData, _ := helper.RsaEncrypt(exampleJSOM)
	
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/rsaExample", strings.NewReader(string(rsaData)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	
	c := e.NewContext(req, rec)
	
	ex := &example.Rsa{}
	
	if assert.NoError(t, ex.RsaPost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, reExampleJSOM, rec.Body.String())
	}
}
