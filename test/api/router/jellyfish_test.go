package router

import (
	"testing"
	"routes/api/router"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

/**
 * JellyFish route test
 */
func TestInitJellyFishRouting(t *testing.T) {
	gin.SetMode(gin.TestMode)
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(cors.Default())
	
	//jellyfish routes
	group := e.Group("")
	router.InitJellyFishRouting(group)
}
