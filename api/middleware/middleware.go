// Package middleware промежуточное API
package middleware

import (
	"net/http"
	"strings"

	"github.com/KnBrBz/snippet-swagger/api"
	"github.com/gin-gonic/gin"
)

// Auth проверка авторизации
func Auth(c *gin.Context) {
	token := extractToken(c.Request)
	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		c.Abort()

		return
	}

	c.Set(api.TokenField, token)

	c.Next()
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	if strArr := strings.Split(bearToken, " "); len(strArr) == 2 { // nolint: gomnd
		return strArr[1]
	}

	return ""
}
