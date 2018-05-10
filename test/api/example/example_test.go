package example

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"routes/api/example"
	"strings"
	"routes/helper"
	"github.com/gin-gonic/gin"
)

var (
	exampleJSOM = `{"id":"111","name":"Jon Snow"}`
	reExampleJSOM = `{"code":"1","result":"{\"id\":\"111\",\"name\":\"Jon Snow\"}"}`
)

/**
 * 測試 GET Info
 */
func TestGet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := &example.Example{}
	
	r.GET("/example/:id/:name", ex.GetInfo)
	req, err := http.NewRequest("GET", "/example/1234/test", nil)
	
	if err != nil {
		t.Fatalf("TestGet NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Host = "CQ1.admin.com"
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("TestGet ServeHTTP Error : %v", err)
	}
}

/**
 *測試 POST
 */

func TestPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := &example.Example{}
	
	r.POST("/example", ex.Post)
	req, err := http.NewRequest("POST", "/example", strings.NewReader(exampleJSOM))
	
	if err != nil {
		t.Fatalf("TestPost NewRequest Error : %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Host = "CQ1.admin.com"
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("TestPost ServeHTTP Error : %v", err)
	}
}

/**
 * 測試 RSA POST
 */

func TestRsa(t *testing.T) {
	rsaData, _ := helper.RsaEncrypt(exampleJSOM)
	gin.SetMode(gin.TestMode)
	
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	
	ex := &example.Rsa{}
	
	r.POST("/rsaExample/Gin", ex.RsaPost)
	req, err := http.NewRequest("POST", "/rsaExample/Gin", strings.NewReader(string(rsaData)))
	
	if err != nil {
		t.Fatalf("TestRsa NewRequest Error : %v", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Host = "CQ1.admin.com"
	
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("TestRsa ServeHTTP Error : %v", err)
	}
}
