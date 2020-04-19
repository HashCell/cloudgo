package api

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/HashCell/golang/cloudgo/pkg/e"
	"github.com/HashCell/golang/cloudgo/models"
	"github.com/HashCell/golang/cloudgo/pkg/util"
	"net/http"
	"github.com/HashCell/golang/cloudgo/pkg/logging"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")

	valid := validation.Validation{}
	authInstance := auth{Username:username, Password:password}
	ok, _ := valid.Valid(&authInstance)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		// check whether the user has existed
		isExist := models.CheckAuth(username, password)
		if isExist {
			// if exists, get token for it
			token, err := util.GenerateToken(username, password)
			logging.Info("hhh")
			if err != nil{
				 code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			// user does not exist
			code = e.ERROR_AUTH
		}
	} else {
		// params error
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}
