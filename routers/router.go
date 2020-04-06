package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/HashCell/golang/cloudgo/pkg/setting"
)

func testHandler(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"message":"test",
	})
}

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	engine.GET("/test", testHandler)

	return engine
}
