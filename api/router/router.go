package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/**
 * 初始化路由
 */
func InitRouting(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	
	//example routes
	exampleGroup := e.Group("/example")
	InitExampleRouting(exampleGroup)
	
	//rsa routes
	rsaGroup := e.Group("/rsaExample")
	InitRsaRouting(rsaGroup)
}
