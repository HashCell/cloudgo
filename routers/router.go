package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/HashCell/golang/cloudgo/pkg/setting"
	"github.com/HashCell/golang/cloudgo/routers/api/v1"
	"github.com/HashCell/golang/cloudgo/routers/api"
	"github.com/HashCell/golang/cloudgo/middleware/jwt"
)

func testHandler(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"message":"test",
	})
}

func InitRouter() *gin.Engine {

	// gin.Engine 继承了　RouterGroup　
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)	// 设置开发debug模式

	// 添加路由处理器
	engine.GET("/auth", api.GetAuth)
	//返回一个RouterGroup
	apiV1 := engine.Group("/api/v1")
	// 对这组api都使用中间件JWT
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		// add routers for article begin 2020.04.14
		// get article list
		apiV1.GET("articles", v1.GetArticles)
		// get specified article
		apiV1.GET("/articles/:id", v1.GetArticle)
		// add article
		apiV1.POST("/articles", v1.AddArticle)
		// update article
		apiV1.PUT("/articles/:id", v1.EditArticle)
		// delete article
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
		// add routers for article end
	}

	return engine
}
