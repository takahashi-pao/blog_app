package middleware

import (
	"fmt"

	"example.com/blog_app/go/internal/handlers/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if auth.Session.Values["userId"] == nil {
			reqUrl := c.Request.URL.Path
			fmt.Print(reqUrl)
			return
		}

		c.Next()
	}
}
