package router

import (
	"github.com/gin-gonic/gin"
	"Stingray/api/example"
)

/**
 * 初始化 Example 路由
 */
func InitExampleRouting(exampleGroup *gin.RouterGroup) {
	//gin.Default()
	
	e := &example.Example{}
	exampleGroup.GET("/:id/:name", e.GetInfo)
	exampleGroup.POST((""), e.Post)
	//exampleGroup.PUT(("/example/:Id"), e.PUTList)
	//exampleGroup.DELETE(("/example/:Id"), e.DELETEList)
}
/**
 * 初始化 RSA 路由
 */
func InitRsaExampleRouting(rsaGroup *gin.RouterGroup) {
	e := &example.Rsa{}
	
	rsaGroup.POST(("/"), e.RsaPost)
}
