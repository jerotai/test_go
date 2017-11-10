package router

import (
	"test_go/api/game"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/**
 * 初始化路由
 */
func InitRouting(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	apiGroup := e.Group("/api")
	apiGroup.GET("/test", game.NewApiGame)

	//e.GET("/test", controller.ApiTest)
}
