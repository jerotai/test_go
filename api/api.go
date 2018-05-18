package api

import (
	"github.com/gin-gonic/gin"
	"Stingray/api/router"
	"github.com/gin-contrib/cors"
)

type APIService struct {
}

func NewApiService() *APIService {
	return &APIService{}
}

/**
 * 啟動API服務
 */
func (s *APIService) Start() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(cors.Default())
	
	ginGroup := g.Group("/example")
	router.InitExampleRouting(ginGroup)
	
	//rsa routes
	exampleGroup := g.Group("/rsaExample")
	router.InitRsaExampleRouting(exampleGroup)
	
	//jellyfish routes
	jellyfishGroup := g.Group("/")
	router.InitJellyFishRouting(jellyfishGroup)
	
	g.Run(":8082")
}
