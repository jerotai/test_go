package api

import (
	"github.com/gin-gonic/gin"
	"routes/api/router"
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
	/*echo := echo.New()
	router.InitRouting(echo)
	echo.Server.Addr = ":8080"
	gracehttp.Serve(echo.Server)
	echo.Close()*/
	//gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	router.InitRouting(g)
	
	g.Run()
}
