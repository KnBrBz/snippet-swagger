package endpoint

import (
	"net/http"

	"github.com/KnBrBz/snippet-swagger/api"
	"github.com/KnBrBz/snippet-swagger/api/middleware"
	"github.com/KnBrBz/snippet-swagger/api/model"
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

// getAll is the handler of all entities
// @Summary Get All
// @Description get all entities
// @Param id path int false  "Entity ID"
// @Tags entities
// @Produce  json
// @Success 200 {object} Posts
// @Router /api/v1/endpoint/{id} [get]
// @Param Authorization header string true "Bearer token"
func getAll(c *gin.Context) {
	code, data := allByFilter(NewFilter(c.Param(ParamID), c.Request.URL.Query()))

	c.JSON(code, gin.H{
		"code": code,
		"data": data,
	})
}

func allByFilter(filter Filter) (int, []Post) {
	data := make([]Post, 0)

	data = append(data, Post{
		Name:  "foo",
		Alias: "bar",
	})

	if len(filter.id) == 0 {
		data = append(data, Post{
			Name:  "bar",
			Alias: "baz",
		})
	}

	return http.StatusOK, data
}

// post is the handler of creating an entity
// @Summary Post
// @Description Post entity
// @Tags entity
// @Accept json
// @Produce  json
// @Param entity body Post true "create entity"
// @Success 200 {object} model.Success
// @Failure 400 {object} model.Success
// @Router /api/v1/endpoint [post]
// @Param Authorization header string true "Bearer token"
func post(c *gin.Context) {
	var t Post

	if err := c.ShouldBindWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, model.Success{Code: http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, model.Success{
		Code: http.StatusOK,
	})
}

// update is the handler of updating entity
// @Summary Update
// @Description update entity
// @Param id path int true  "Entity ID"
// @Tags entity
// @Accept json
// @Produce  json
// @Param entity body Post true "update entity"
// @Success 200 {object} model.Success
// @Failure 400 {object} model.Success
// @Router /api/v1/endpoint/{id} [put]
// @Param Authorization header string true "Bearer token"
func update(c *gin.Context) {
	var t Post

	if err := c.ShouldBindWith(&t, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, model.Success{Code: http.StatusBadRequest})
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

// delete is the handler of deleting entity
// @Summary Delete
// @Description delete entity
// @Param id path int true  "Entity ID"
// @Tags entity
// @Produce  json
// @Success 200 {object} model.Success
// @Failure 400 {object} model.Success
// @Router /api/v1/endpoint/{id} [delete]
// @Param Authorization header string true "Bearer token"
func del(c *gin.Context) {
	c.JSON(deleteByID(c.Param(ParamID)))
}

func deleteByID(id string) (int, model.Success) {
	if len(id) == 0 {
		return http.StatusBadRequest, model.Success{Code: http.StatusBadRequest}
	}

	return http.StatusOK, model.Success{Code: http.StatusOK}
}
