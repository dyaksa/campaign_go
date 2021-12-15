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

	token, err := h.authService.GenerateToken(existUser.ID)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("token error", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
	}

	responseUserFormatter := user.FormatUser(existUser, token)
	responseJson := helper.APIResponse("success login", http.StatusOK, "success", responseUserFormatter)
	c.JSON(http.StatusOK, responseJson)
}

func (h *userHandler) CheckEmail(c *gin.Context) {
	input := user.CheckEmailInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("unprocessable entity", http.StatusUnprocessableEntity, "errors", data)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}
	exist, err := h.service.EmailChecker(input.Email)
	if err != nil {
		data := gin.H{"error": err.Error()}
		responseJSON := helper.APIResponse("bad request", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}

	if exist.ID != 0 {
		data := gin.H{"is_available": false}
		responseJSON := helper.APIResponse("email already in use", http.StatusConflict, "errors", data)
		c.JSON(http.StatusConflict, responseJSON)
		return
	}
	data := gin.H{"is_available": true}
	responseJSON := helper.APIResponse("email available", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, responseJSON)
}

func (h *userHandler) UpdateAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		responseJSON := helper.APIResponse("fail upload avatar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	user := c.MustGet("user").(user.User)
	path := fmt.Sprintf("images/%d-%s", user.ID, file.Filename)
	c.SaveUploadedFile(file, path)
	_, err = h.service.UpdateAvatar(user.ID, path)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("errors update avatar", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
	}
	data := gin.H{"is_uploaded": true}
	responseJSON := helper.APIResponse("success update avatar", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, responseJSON)
}
