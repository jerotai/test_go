package apicurl

import (
	"github.com/gin-gonic/gin"
	"Stingray/trait"
	"Stingray/helper"
	"encoding/base64"
	"net/http"
)

/**
 * get new rsa key by token
 */
func (s *apiServiceCurl) NewRsaPubKey (g *gin.Context) {
	ApiToken := trait.GetToken(g)
	
	pub, pri := helper.CreateRsaKey()
	
	trait.SaveLoginRsaKey(ApiToken, pub, pri)
	
	var rsaKey = make(map[string]string)
	rsaKey["key"] = base64.StdEncoding.EncodeToString([]byte(pub))
	
	reBody := &trait.ApiResponse{}
	reBody.Code = "1"
	reBody.Result = rsaKey
	g.JSON(http.StatusOK, reBody)
}