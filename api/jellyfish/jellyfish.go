package jellyfish

import (
	"Stingray/helper"
	"github.com/gin-gonic/gin"
)

type (
	// JellyFish 建構子
	JellyFish struct {
		//SendCurl *helper.CurlBase
		SendCurl *helper.CurlBase
		ApiConf struct {
			dto interface{}
			apiRequestUrl string
		}
	}
)

func New() *JellyFish {
	return &JellyFish{}
}

func (j *JellyFish)GetToken(g *gin.Context) string {
	return g.Request.Header.Get("Api-Token")
}