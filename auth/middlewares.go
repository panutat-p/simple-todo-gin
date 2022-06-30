package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

// ValidateAccessToken act as auth guard
func ValidateAccessToken(signature []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(bearerToken, "Bearer ")

		// higher order function for looking up the key
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return signature, nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized) // break middleware chain
			return
		}
		fmt.Println("token:", token)
		c.Next() // continue to middleware chain
	}
}
