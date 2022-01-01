package handler

import (
	"campaignproject/helper"
	"campaignproject/transaction"
	"campaignproject/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionsHandler struct {
	service transaction.Service
}

func NewTransactionsHandler(service transaction.Service) *transactionsHandler {
	return &transactionsHandler{service: service}
}

func (h *transactionsHandler) GetByCampaignID(c *gin.Context) {
	var input transaction.GetCampaignID
	var pagination helper.Pagination
	currentUser := c.MustGet("user").(user.User)
	err := c.Bind(&pagination)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		data := gin.H{"errors": errors}
		responseJSON := helper.APIResponse("failed get transactions", http.StatusUnprocessableEntity, "errors", data)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}

	err = c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		data := gin.H{"errors": errors}
		responseJSON := helper.APIResponse("failed get transaction", http.StatusUnprocessableEntity, "errors", data)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}
	input.User = currentUser
	transactions, err := h.service.GetTransacationsByCampaignID(input, pagination)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("error get transactions", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	responseJSON := helper.APIResponse("success get transactions by campaign", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, responseJSON)
}

func (h *transactionsHandler) GetTransactionsByUserID(c *gin.Context) {
	currentUser := c.MustGet("user").(user.User)
	pagination := helper.Pagination{}
	c.Bind(&pagination)
	transactions, err := h.service.GetTransactionsByUserID(currentUser.ID, pagination)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("failed get transactions by user id", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	responseJSON := helper.APIResponse("success get transactions user", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, responseJSON)
}

func (h *transactionsHandler) CreateTransaction(c *gin.Context) {
	input := transaction.InputTransaction{}
	currentUser := c.MustGet("user").(user.User)
	input.UserID = currentUser.ID
	input.Name = currentUser.Name
	input.Email = currentUser.Email
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		data := gin.H{"errors": errors}
		responseJSON := helper.APIResponse("failed create transaction", http.StatusUnprocessableEntity, "errors", data)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}
	transaction, err := h.service.SaveTransaction(input)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("failed create transaction", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	responseJSON := helper.APIResponse("success create transaction", http.StatusCreated, "success", transaction)
	c.JSON(http.StatusCreated, responseJSON)
}
