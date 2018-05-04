package router

import (
	"testing"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"routes/api/router"
)

/**
 * Example route test
 */
func TestInitExampleRouting(t *testing.T) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	
	//example routes
	exampleGroup := e.Group("/example")
	router.InitExampleRouting(exampleGroup)
}

/**
 * rsa route test
 */
func TestInitRsaRouting(t *testing.T) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	
	//rsa routes
	rsaGroup := e.Group("/rsaExample")
	router.InitRsaRouting(rsaGroup)
}