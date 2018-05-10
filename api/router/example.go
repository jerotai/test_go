package router

import (
	"routes/api/example"
	"github.com/gin-gonic/gin"
)

/**
 * 初始化 Example 路由
 */
/*func InitExampleRouting(exampleGroup *echo.Group) {
	e := &example.Example{}
	
	exampleGroup.GET(("/:id"), e.GetInfo)
	exampleGroup.GET(("/:id/:name"), e.GetInfo)
	exampleGroup.POST(("/"), e.Post)
	//exampleGroup.POST(("/rsa/"), e.RsaPostList)
	//exampleGroup.PUT(("/example/:Id"), e.PUTList)
	//exampleGroup.DELETE(("/example/:Id"), e.DELETEList)
}*/

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
/*func InitRsaRouting(rsaGroup *echo.Group) {
	e := &example.Rsa{}
	
	rsaGroup.POST(("/"), e.RsaPost)
}*/

func InitRsaRouting(rsaGroup *gin.RouterGroup) {
	e := &example.Rsa{}
	
	rsaGroup.POST(("/"), e.RsaPost)
}
