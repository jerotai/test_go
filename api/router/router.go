package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

/**
 * 初始化路由
 */
/*func InitRouting(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	
	//example routes
	exampleGroup := e.Group("/example")
	InitExampleRouting(exampleGroup)
	
	//rsa routes
	rsaGroup := e.Group("/rsaExample")
	InitRsaRouting(rsaGroup)
}*/

func InitRouting(g *gin.Engine) {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(cors.Default())
	
	ginGroup := g.Group("/example")
	InitExampleRouting(ginGroup)
	
	//rsa routes
	rsaGroup := g.Group("/rsaExample")
	InitRsaRouting(rsaGroup)
	
	//jellyfish routes
	jellyfishGroup := g.Group("/")
	InitJellyFishRouting(jellyfishGroup)
}
