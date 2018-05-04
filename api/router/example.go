package router

import (
	"github.com/labstack/echo"
	"routes/api/example"
)

/**
 * 初始化路由
 */
func InitExampleRouting(exampleGroup *echo.Group) {
	e := &example.Example{}
	
	exampleGroup.GET(("/:id"), e.GetInfo)
	exampleGroup.GET(("/:id/:name"), e.GetInfo)
	exampleGroup.POST(("/"), e.Post)
	//exampleGroup.POST(("/rsa/"), e.RsaPostList)
	//exampleGroup.PUT(("/example/:Id"), e.PUTList)
	//exampleGroup.DELETE(("/example/:Id"), e.DELETEList)
}

/**
 * 初始化路由
 */
func InitRsaRouting(rsaGroup *echo.Group) {
	e := &example.Rsa{}
	
	rsaGroup.POST(("/"), e.RsaPost)
	
}
