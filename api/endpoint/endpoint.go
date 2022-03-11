package endpoint

import (
	"net/http"

	"github.com/KnBrBz/snippet-swagger/api"
	"github.com/KnBrBz/snippet-swagger/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	APIVersion = api.APIVersion
	APIPath    = api.APIPath

	ParamID = "id"
)

func Init(r *gin.Engine) {
	v := r.Group(APIPath + APIVersion + "/endpoint")

	v.GET("/", middleware.Auth, getAll)          // ALL
	v.GET("/:"+ParamID, middleware.Auth, getAll) // SINGLE
	v.POST("/", middleware.Auth, post)
	v.PUT("/:"+ParamID, middleware.Auth, update)
	v.DELETE("/:"+ParamID, middleware.Auth, del)
}

func getAll(c *gin.Context) {
	code, data := allByFilter(NewFilter(c.Param(ParamID), c.Request.URL.Query()))

	c.JSON(code, gin.H{
		"code": code,
		"data": data,
	})
}

func allByFilter(filter Filter) (int, []gin.H) {
	data := make([]gin.H, 0)

	data = append(data, gin.H{
		"name":  "foo",
		"alias": "bar",
	})

	if len(filter.id) == 0 {
		data = append(data, gin.H{
			"name":  "bar",
			"alias": "baz",
		})
	}

	return http.StatusOK, data
}

func post(c *gin.Context) {
	var t Post

	if err := c.ShouldBindWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func update(c *gin.Context) {
	var t Post

	if err := c.ShouldBindWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest})
		return
	}

	c.JSON(updateByID(c.Param(ParamID), t))
}

func updateByID(id string, t Post) (int, gin.H) {
	if len(t.Alias) == 0 {
		return http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		}
	}

	if len(id) == 0 {
		return http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		}
	}

	return http.StatusOK, gin.H{
		"code": http.StatusOK,
	}
}

func del(c *gin.Context) {
	c.JSON(deleteByID(c.Param(ParamID)))
}

func deleteByID(id string) (int, gin.H) {
	if len(id) == 0 {
		return http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		}
	}

	return http.StatusOK, gin.H{
		"code": http.StatusOK,
	}
}
