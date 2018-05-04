package example

import (
	"github.com/labstack/echo"
	"net/http"
	"routes/core/dto"
	"routes/api/apiparse"
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
//noinspection GoTypesCompatibility
func (ex *Example) GetInfo(c echo.Context) error {
	req := dto.ExampleList{}
	//req.Id = c.Param("id")
	//req.Name = c.Param("name")
	apiparse.GetJsonDataParse(c, &req)
	return c.JSON(http.StatusOK, apiparse.ApiResponse(req))
}

// api post 範例
func (ex *Example) Post(c echo.Context) error {
	req := dto.ExampleList{}
	apiparse.PostJsonDataParse(c, &req)
	return c.JSON(http.StatusOK, apiparse.ApiResponse(req))
}

// api rsa post 範例
func (ex *Rsa) RsaPost(c echo.Context) error {
	req := dto.ExampleList{}
	apiparse.PostJsonRsaDataParse(c, &req)
	return c.JSON(http.StatusOK, apiparse.ApiResponse(req))
}