package router

import (
	//"test_go/api/controller/api1"
	//"test_go/api/controller/api2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/**
 * 初始化路由
 */
func InitRouting(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//api2Group := e.Group("/api2")
	//api2Group.GET("/test", api2.ApiTest)

	//e.GET("/test", controller.ApiTest)
}
