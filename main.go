package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/panutat-p/simeple-todo-gin/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load ENV")
	}
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect ElephantSQL")
	}
	fmt.Println("Connected to ElephantSQL")

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate ElephantSQL")
	}

	userHandler := user.NewHandler(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	r.GET("/user", userHandler.GetFirstUser)

	err = r.Run() // block
	if err != nil {
		log.Println("server crashed")
		return
	}
}
