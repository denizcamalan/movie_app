package middleware

import (
	"net/http"

	"github.com/denizcamalan/movie_app/repo"
	"github.com/gin-gonic/gin"
)

var jwt_controller = repo.NewOperatorModel().JWT_Operator()

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt_controller.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}