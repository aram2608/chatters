package middleware

import (
	"chatters-REST/config"
	"chatters-REST/controllers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string

		// We get the auth header
		auth := c.GetHeader("Authorization")
		if auth != "" && strings.HasPrefix(auth, "Bearer ") {
			tokenStr = strings.TrimPrefix(auth, "Bearer ")
		}

		// We fail out if the token string is empty
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		// We try to parse the claims
		token, err := jwt.ParseWithClaims(
			tokenStr,
			&controllers.Claims{},
			func(t *jwt.Token) (any, error) {
				return config.JwtSecret, nil
			})

		// We abort out if we catch an error or the token is bad
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// If the claim was not okay we abort out
		claims, ok := token.Claims.(*controllers.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		// We set user info for handlers
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		// We go to the next handler
		c.Next()
	}
}
