package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	Name string
}

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
	fmt.Println("DB connection is established")

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate ElephantSQL")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		var u User
		db.First(&u)
		c.JSON(http.StatusOK, u)
	})

	err = r.Run() // block
	if err != nil {
		log.Println("server crashed")
		return
	}
}
