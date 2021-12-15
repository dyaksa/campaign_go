package main

import (
	"campaignproject/auth"
	"campaignproject/campaign"
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

	//* user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	middleware := middleware.NewMiddleware(authService, userService)
	authMiddleware := middleware.AuthUser()

	//* campaign
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	app := router.Group("/api/v1")
	app.POST("/register", userHandler.RegisterUser)
	app.POST("/login", userHandler.Login)
	app.PATCH("/upload", authMiddleware, userHandler.UpdateAvatar)

	app.GET("/campaign", campaignHandler.FindAllCampaign)
	app.POST("/campaign", authMiddleware, campaignHandler.InputInsertCampaign)
	app.GET("/campaign/:slug", campaignHandler.DetailBySlug)
	app.GET("/auth/user/campaign", authMiddleware, campaignHandler.UserHaveCampaigns)
	router.Run(":8080")
}
