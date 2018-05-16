package jellyfish

import (
	"Stingray/helper"
)

type (
	// JellyFish 建構子
	JellyFish struct {
		//SendCurl *helper.CurlBase
		SendCurl *helper.CurlBase
	}
)

func New() *JellyFish {
	return &JellyFish{}
}