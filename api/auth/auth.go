package auth

import (
	"log"
	"net/http"

	"github.com/KnBrBz/snippet-swagger/api"
	"github.com/KnBrBz/snippet-swagger/api/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
)

const packageTitle = "auth: "

func Init(r *gin.Engine) {
	v := r.Group(api.APIPath + api.APIVersion + "/auth")

	v.POST("/", getToken)
}

// getToken is the handler of token generation
// @Summary Token
// @Description get token
// @Param authData body Auth true  "Login, Password"
// @Tags authorization
// @Produce  json
// @Success 200 {object} AuthSuccess
// @Failure 400 {object} model.Success
// @Router /api/v1/auth [post]
func getToken(c *gin.Context) {
	const funcTitle = packageTitle + "getToken"

	var t Auth

	if err := c.ShouldBindWith(&t, binding.JSON); err != nil {
		log.Println(errors.Wrap(err, funcTitle))
		c.JSON(http.StatusBadRequest, model.Success{Code: http.StatusBadRequest})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": AuthSuccess{
			Token: "some_token",
		},
	})
}
