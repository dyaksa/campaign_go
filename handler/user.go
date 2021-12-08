package handler

import (
	"campaignproject/auth"
	"campaignproject/helper"
	"campaignproject/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service: service, authService: authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		responseError := gin.H{"errors": errors}
		response := helper.APIResponse("register user failed", http.StatusBadRequest, "failed", responseError)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.service.RegisterInput(input)
	if err != nil {
		response := helper.APIResponse("register user failed", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userFormatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("user success has been registered", http.StatusOK, "success", userFormatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		inputError := helper.FormatterErroValidation(err)
		responseError := gin.H{"errors": inputError}
		responseJson := helper.APIResponse("failed login", http.StatusUnprocessableEntity, "failed", responseError)
		c.JSON(http.StatusUnprocessableEntity, responseJson)
		return
	}

	existUser, err := h.service.Login(input)
	if err != nil {
		responseError := gin.H{"errors": err.Error()}
		responseJson := helper.APIResponse("failed login", http.StatusUnprocessableEntity, "failed", responseError)
		c.JSON(http.StatusUnprocessableEntity, responseJson)
		return
	}

	token, err := h.authService.GenerateToken(existUser.ID, existUser.Email, existUser.Name, existUser.Occupation)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("token error", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
	}

	responseUserFormatter := user.FormatUser(existUser, token)
	responseJson := helper.APIResponse("success login", http.StatusOK, "success", responseUserFormatter)
	c.JSON(http.StatusOK, responseJson)
}

func (h *userHandler) UpdateAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		responseJSON := helper.APIResponse("fail upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	userID := 12
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	c.SaveUploadedFile(file, path)
	_, err = h.service.UpdateAvatar(userID, path)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("errors update avatar", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
	}
	data := gin.H{"is_uploaded": true}
	responseJSON := helper.APIResponse("success update avatar", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, responseJSON)
}
