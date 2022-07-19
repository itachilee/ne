package middleware

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		gin.BasicAuth(gin.Accounts{
			"user1": "love",
			"user2": "god",
			"user3": "sex",
		})
	}
}
