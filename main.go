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
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// dsn := "root:@tcp(127.0.0.1:3306)/campaign_project?charset=utf8mb4&parseTime=True&loc=Local"
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

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	app := router.Group("/api/v1")
	app.GET("/", func(c *gin.Context) {
		data := gin.H{"message": "welcome to campaign server"}
		c.JSON(http.StatusOK, data)
	})
	app.POST("/register", userHandler.RegisterUser)
	app.POST("/login", userHandler.Login)
	app.POST("/email_checker", userHandler.CheckEmail)
	app.POST("/upload", authMiddleware, userHandler.UpdateAvatar)
	app.GET("/user/fetch", authMiddleware, userHandler.FetchUser)

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

	router.Run(":" + os.Getenv("PORT"))
}
