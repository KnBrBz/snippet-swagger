package auth

import (
	"log"
	"net/http"

	"github.com/KnBrBz/snippet-swagger/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
)

const packageTitle = "auth: "

func Init(r *gin.Engine) {
	v := r.Group(api.APIPath + api.APIVersion + "/auth")

	v.POST("/", getToken)
}

func getToken(c *gin.Context) {
	const funcTitle = packageTitle + "getToken"

	var t Auth

	if err := c.ShouldBindWith(&t, binding.JSON); err != nil {
		log.Println(errors.Wrap(err, funcTitle))
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"token": "some_token",
		},
	})
}
