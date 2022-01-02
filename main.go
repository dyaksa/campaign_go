package main

import (
	"campaignproject/auth"
	"campaignproject/campaign"
	"campaignproject/handler"
	"campaignproject/middleware"
	"campaignproject/payment"
	"campaignproject/transaction"
	"campaignproject/user"
	"fmt"

	"github.com/gin-contrib/cors"
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
	paymentService := payment.NewPaymantService()

	//* campaign
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	//* transactions
	transactionsRepository := transaction.NewRepository(db)
	transactionsService := transaction.NewService(transactionsRepository, campaignRepository, paymentService)
	transactionsHandler := handler.NewTransactionsHandler(transactionsService)

	router := gin.Default()
	config := cors.DefaultConfig()
	router.Use(cors.New(config))
	router.Static("/images", "./images")
	app := router.Group("/api/v1")
	app.POST("/register", userHandler.RegisterUser)
	app.POST("/login", userHandler.Login)
	app.POST("/email_checker", userHandler.CheckEmail)
	app.POST("/upload", authMiddleware, userHandler.UpdateAvatar)

	app.GET("/campaign", campaignHandler.FindAllCampaign)
	app.POST("/campaign", authMiddleware, campaignHandler.InputInsertCampaign)
	app.GET("/campaign/:slug", campaignHandler.DetailBySlug)
	app.PUT("/campaign/:id", authMiddleware, campaignHandler.UpdateCampaign)
	app.GET("/auth/user/campaign", authMiddleware, campaignHandler.UserHaveCampaigns)
	app.POST("/campaign/upload/images", authMiddleware, campaignHandler.UploadCampaignImages)

	app.GET("/transaction/campaign/:id", authMiddleware, transactionsHandler.GetByCampaignID)
	app.GET("/transactions", authMiddleware, transactionsHandler.GetTransactionsByUserID)
	app.POST("/transaction", authMiddleware, transactionsHandler.CreateTransaction)
	app.POST("/transaction/notification", transactionsHandler.PaymentNotification)

	router.Run(":8080")
}
