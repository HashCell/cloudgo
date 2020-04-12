package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/HashCell/golang/cloudgo/pkg/status"
	"github.com/HashCell/golang/cloudgo/models"
	"github.com/HashCell/golang/cloudgo/pkg/util"
	"github.com/HashCell/golang/cloudgo/pkg/setting"
	"net/http"
	"github.com/astaxie/beego/validation"
	"log"
)


// get tag list of articles
func GetTags(ctx *gin.Context) {

	// get query string and distract them out
	// organize them as query condition for db
	name := ctx.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	// put all querying variables for db query into mao
	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()	// transfer into int
		maps["state"] = state
	}

	code := status.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	ctx.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg":status.GetMsg(code),
		"data":data,
	})
}

// add an article tag
func AddTag(ctx *gin.Context) {
	name := ctx.Query("name")
	state, _ := com.StrTo(ctx.DefaultQuery("state","0")).Int()
	createdBy := ctx.Query("created_by")

	log.Println("[DEBUG]:", name, state, createdBy)
	// 使用beego的表单验证包
	valid := validation.Validation{}
	valid.Required(name, "name").Message("name cannot be null")
	valid.Required(name, "created_by").Message("creator cannot be null")
	valid.MaxSize(name, 100, "created_by").Message("max length is 100")
	valid.MaxSize(name, 100, "name").Message("max length of name is 100")
	valid.Range(state, 0, 1, "state").Message("state should be 0 - 1")

	code := status.INVALID_PARAMS
	if !valid.HasErrors() {
		if ! models.ExistTagName(name) {
			code = status.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = status.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range  valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg":status.GetMsg(code),
		"data":make(map[string]string),
	})
}

// edit tags of some article
func EditTag(ctx *gin.Context) {
	id, _ := com.StrTo(ctx.Param("id")).Int()
	name := ctx.Query("name")
	modifiedBy := ctx.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state, _ := com.StrTo(arg).Int()
		valid.Range(state, 0, 1, "state").Message("state must 0 - 1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := status.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = status.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = status.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : status.GetMsg(code),
		"data" : make(map[string]string),
	})
}

// delete tags of some article
func DeleteTag(ctx *gin.Context) {
	id, _ := com.StrTo(ctx.Param("id")).Int()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := status.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = status.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = status.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : status.GetMsg(code),
		"data" : make(map[string]string),
	})
}


