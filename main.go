package main

import (
	"campaignproject/auth"
	"campaignproject/handler"
	"campaignproject/middleware"
	"campaignproject/user"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/campaign_project?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("success connect db")
	authService := auth.NewAuthService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	app := router.Group("/api/v1")
	app.POST("/users", userHandler.RegisterUser)
	app.POST("/login", userHandler.Login)
	app.PATCH("/upload", middleware.AuthUser(authService), userHandler.UpdateAvatar)
	router.Run()

}
