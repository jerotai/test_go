package jellyfish

import (
	"routes/helper"
)

type (
	// JellyFish 建構子
	JellyFish struct {
		SendCurl *helper.CurlBase
	}
)

func New() *JellyFish {
	return &JellyFish{}
}