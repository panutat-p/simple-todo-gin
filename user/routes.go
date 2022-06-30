package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GetFirstUser(c *gin.Context) {
	var u User
	h.db.First(&u)
	s, _ := json.MarshalIndent(u, "", "\t")
	fmt.Println(string(s))
	c.JSON(http.StatusOK, u)
}
