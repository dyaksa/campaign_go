package handler

import (
	"campaignproject/helper"
	"campaignproject/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service: service}
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

	responseUserFormatter := user.FormatUser(existUser, "tokentokentoken")
	responseJson := helper.APIResponse("success login", http.StatusOK, "success", responseUserFormatter)
	c.JSON(http.StatusOK, responseJson)
}
