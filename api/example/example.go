package example

import (
	"net/http"
	"routes/core/dto"
	"routes/api/apiparse"
	"github.com/gin-gonic/gin"
	"fmt"
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
	req.SiteCode = apiparse.SiteCodeParse(g.Request)
	fmt.Println(req)
	g.JSON(200, req)
}

/**
 * api rsa post 範例
 */
func (ex *Rsa) RsaPost(g *gin.Context) {
	req := dto.ExampleList{}
	apiparse.PostRsaDataParse(g, &req)
	g.JSON(200, req)
}

/**
 * api Post 範例
 */
func (ex *Example) Post(g *gin.Context) {
	req := dto.ExampleList{}
	apiparse.PostDataParse(g, &req)
	fmt.Println(req)
	g.JSON(200, req)
}

/**
 * echo api get 範例
 */
/*
func (ex *Example) GetInfoByEcho(c echo.Context) error {
	req := dto.ExampleList{}
	//req.Id = c.Param("id")
	//req.Name = c.Param("name")
	apiparse.GetJsonDataParse(c, &req)
	return c.JSON(http.StatusOK, apiparse.ApiResponse(req))
}
*/

/**
 * echo api post 範例
 */
/*
func (ex *Example) PostByEcho(c echo.Context) error {
	req := dto.ExampleList{}
	apiparse.PostDataParse(c, &req)
	return c.JSON(http.StatusOK, apiparse.ApiResponse(req))
}
*/

/**
 * echo api rsa post 範例
 */
/*
func (ex *Rsa) RsaPostByEcho(c echo.Context) error {
	req := dto.ExampleList{}
	apiparse.PostRsaDataParse(c, &req)
	return c.JSON(http.StatusOK, apiparse.ApiResponse(req))
}
*/

