package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/HashCell/golang/cloudgo/pkg/e"
	"github.com/HashCell/golang/cloudgo/pkg/util"
	"time"
	"net/http"
)

// 返回一个路由处理函数
func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := context.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			// token解析错误
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				// token过期
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":code,
				"msg":e.GetMsg(code),
				"data":data,
			})

			context.Abort()
			return
		}

		// 中间件，调用链
		context.Next()
	}
}