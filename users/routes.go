package users

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	Db *gorm.DB
}

func (h *UserHandler) GetFirstUser(c *gin.Context) {
	var u User
	h.Db.First(&u)
	s, _ := json.MarshalIndent(u, "", "\t")
	fmt.Println(string(s))
	c.JSON(http.StatusOK, u)
}
