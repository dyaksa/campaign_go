package main

import (
	"campaignproject/campaign"
	"campaignproject/handler"
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
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()

	app := router.Group("/api/v1")
	app.POST("/users", userHandler.RegisterUser)
	app.POST("/campaign", campaignHandler.InputInsertCampaign)
	app.GET("/campaign", campaignHandler.GetCampaign)
	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Email = "diasnour0395@gmail.com"
	// userInput.Name = "Dyaksa Jauhruddin"
	// userInput.Occupation = "programmer"
	// userInput.Password = "secret"
	// userInput.Role = "user"

	// userService.RegisterInput(userInput)

	// userHandler := handler.NewUserHandler(userService)

}
