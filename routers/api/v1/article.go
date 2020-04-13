package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/HashCell/golang/cloudgo/pkg/status"
	"github.com/HashCell/golang/cloudgo/models"
	"log"
	"net/http"
	"github.com/HashCell/golang/cloudgo/pkg/util"
	"github.com/HashCell/golang/cloudgo/pkg/setting"
)

// get single article
func GetArticle(ctx *gin.Context) {
	id, _ := com.StrTo(ctx.Param("id")).Int()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID must > 1")

	code := status.INVALID_PARAMS
	var data interface{}
	if!valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = status.SUCCESS
		} else {
			code = status.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Println("[DEBUG]: ", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": status.GetMsg(code),
		"data": data,
	})
}

func GetArticles(ctx *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("state only allow 0 or 1")
	}

	var tagId = -1
	if arg := ctx.Query("tag_id"); arg != "" {
		tagId, _ = com.StrTo(arg).Int()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("tagId must > 1")
	}

	code := status.INVALID_PARAMS
	if !valid.HasErrors() {
		code = status.SUCCESS
		data["lists" ] = models.GetArticles(util.GetPage(ctx), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Println("[DEBUG]: ", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": status.GetMsg(code),
		"data": data,
	})
}

func AddArticle(ctx *gin.Context) {
	tagId, _ := com.StrTo(ctx.Query("tag_id")).Int()
	title := ctx.Query("title")
	desc := ctx.Query("desc")
	content := ctx.Query("content")
	createdBy := ctx.Query("created_by")
	state, _ := com.StrTo(ctx.DefaultQuery("state", "0")).Int()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("tag ID must > 0")
	valid.Required(title, "title").Message("title cannot be null")
	valid.Required(desc, "desc").Message("desc of article cannot be null")
	valid.Required(content, "content").Message("content cannot be null")
	valid.Required(createdBy, "created_by").Message("author cannot be null")
	valid.Range(state, 0, 1, "state").Message("state only allowed for 0 or 1")

	code := status.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface {})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = status.SUCCESS
		} else {
			code = status.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Println("[DEBUG]:", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : status.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

func EditArticle(ctx *gin.Context) {
	valid := validation.Validation{}

	id, _ := com.StrTo(ctx.Param("id")).Int()
	tagId, _ := com.StrTo(ctx.Query("tag_id")).Int()
	title := ctx.Query("title")
	desc := ctx.Query("desc")
	content := ctx.Query("content")
	modifiedBy := ctx.Query("modified_by")

	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		valid.Range(state, 0, 1, "state").Message("state only allowed for 0 or 1")
	}

	valid.Min(id, 1, "id").Message("ID must > 0")
	valid.MaxSize(title, 100, "title").Message("max length of title < 100 char")
	valid.MaxSize(desc, 255, "desc").Message("max lenght of desc < 255 char")
	valid.MaxSize(content, 65535, "content").Message("max length of content < 65535")
	valid.Required(modifiedBy, "modified_by").Message("modified_by cannot be null")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("max length of modified_by < 100")

	code := status.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			if models.ExistTagByID(tagId) {
				data := make(map[string]interface {})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = status.SUCCESS
			} else {
				code = status.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = status.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Println("[DEBUG]:", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : status.GetMsg(code),
		"data" : make(map[string]string),
	})
}

func DeleteArticle(ctx *gin.Context) {
	id, _ := com.StrTo(ctx.Param("id")).Int()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID must > 0")

	code := status.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = status.SUCCESS
		} else {
			code = status.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Println("[DEBUG]:", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : status.GetMsg(code),
		"data" : make(map[string]string),
	})
}


