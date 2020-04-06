package util

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/HashCell/golang/cloudgo/pkg/setting"
)

func GetPage(ctx *gin.Context) int {
	result := 0
	page, _ := com.StrTo(ctx.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
