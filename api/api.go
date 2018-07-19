package api

import (
	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	"Stingray/api/router"
)

type APIService struct {
}

func NewApiService() *APIService {
	apiService := &APIService{}
	return apiService
}

/**
 * 啟動API服務
 */
func (s *APIService) Start() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(limit.MaxAllowed(20000))
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	
	corsOptions := Options{}
	//corsOptions.AllowOrigins = []string{"http://moa.cqcp.corp"}
	corsOptions.AllowCredentials = true
	g.Use(Middleware(corsOptions))
	
	//g.Use(cors.Default())
	
	//jellyfish routes
	jellyfishGroup := g.Group("/")
	router.InitJellyFishRouting(jellyfishGroup)
	
	//whale routes
	whaleGroup := g.Group("/")
	router.InitWhaleRouting(whaleGroup)
	
	//whitebait routes
	whitebaitGroup := g.Group("/")
	router.InitWhitebaitRouting(whitebaitGroup)
	
	//lophiiformes routes
	lophiiformesGroup := g.Group("/")
	router.InitLophiiformesRouting(lophiiformesGroup)
	
	//octopus routes
	octopusGroup := g.Group("/")
	router.InitOctopusRouting(octopusGroup)
	
	//lobster routes
	lobsterGroup := g.Group("/")
	router.InitLobsterRouting(lobsterGroup)
	
	//electricity routes
	electricityGroup := g.Group("/")
	router.InitElectricityRouting(electricityGroup)
	
	//stingray routes
	stingrayGroup := g.Group("/")
	router.InitStingrayRouting(stingrayGroup)
	
	g.Run(":8082")
}
