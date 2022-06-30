package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/panutat-p/simeple-todo-gin/auth"
	"github.com/panutat-p/simeple-todo-gin/todo"
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

	err = db.AutoMigrate(&user.User{}, &todo.Todo{})
	if err != nil {
		panic("Failed to migrate ElephantSQL")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	authHandler := auth.NewHandler([]byte(jwtSecret))
	userHandler := user.NewHandler(db)
	todoHandler := todo.NewHandler(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	r.POST("/auth/login", authHandler.SignAccessToken)

	r.GET("/user", userHandler.GetFirstUser)

	g1 := r.Group("", auth.ValidateAccessToken([]byte(jwtSecret)))
	g1.GET("/todos", todoHandler.GetAllTasks)
	g1.POST("/todo/new", todoHandler.NewTask)

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))) // block
	if err != nil {
		log.Println("🟥 Cannot start web server")
		return
	}
}
