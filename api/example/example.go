package example

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"Stingray/core/dto"
	"Stingray/api/apiparse"
)

type (
	// Example 建構子
	Example struct {
		http *http.Request
	}
	
	// RSA 建構子
	Rsa struct {
	}
)

/**
 * api get 範例
 */
func (ex *Example) GetInfo(g *gin.Context) {
	req := dto.ExampleList{}
	//parse api get params
	apiparse.GetDataParse(g, &req)
	
	//get site code
	req.StationCode = apiparse.StationCodeParse(g.Request)
	fmt.Println("GetInfo", req)
	g.JSON(200, req)
}

/**
 * api rsa post 範例
 */
func (ex *Rsa) RsaPost(g *gin.Context) {
	req := dto.ExampleList{}
	apiparse.RsaDataParse(g, &req)
	fmt.Println("RsaPost", req)
	g.JSON(200, req)
}

/**
 * api Post 範例
 */
func (ex *Example) Post(g *gin.Context) {
	req := dto.ExampleList{}
	apiparse.PostDataParse(g, &req)
	fmt.Println("Post", req)
	g.JSON(200, req)
}
