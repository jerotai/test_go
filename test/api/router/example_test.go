package router

import (
	"testing"
	"Stingray/api/router"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

/**
 * Example route test
 */
func TestInitExampleRouting(t *testing.T) {
	gin.SetMode(gin.TestMode)
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(cors.Default())
	
	//example routes
	exampleGroup := e.Group("/example")
	router.InitExampleRouting(exampleGroup)
}

/**
 * rsa route test
 */
func TestInitRsaRouting(t *testing.T) {
	gin.SetMode(gin.TestMode)
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(cors.Default())
	
	//rsa routes
	rsaGroup := e.Group("/rsaExample")
	router.InitRsaRouting(rsaGroup)
}