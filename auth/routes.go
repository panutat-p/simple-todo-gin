package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	secret []byte
}

func NewHandler(secret []byte) *Handler {
	return &Handler{
		secret: secret,
	}
}

// SignAccessToken get encoded format of JWT token string
func (h Handler) SignAccessToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Local().Add(time.Hour * time.Duration(12)).Unix(),
	})
	tokenString, err := token.SignedString(h.secret)
	if err != nil {
		log.Println("🟥 Failed to sign a token")
		log.Println(err)
		return
	}
	s, _ := json.MarshalIndent(map[string]string{"accessToken": tokenString}, "", "  ")
	fmt.Println(string(s))
	c.JSON(http.StatusOK, gin.H{
		"accessToken": tokenString,
	})
}
