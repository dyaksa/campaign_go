package handler

import (
	"campaignproject/helper"
	"campaignproject/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service transaction.Service
}

func NewTransactionsHandler(service transaction.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetByCampaignID(c *gin.Context) {
	var input transaction.GetCampaignID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		data := gin.H{"errors": errors}
		responseJSON := helper.APIResponse("failed get transaction", http.StatusUnprocessableEntity, "failed", data)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}
	transactions, err := h.service.GetTransacationsByCampaignID(input.ID)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("error get transactions", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	responseJSON := helper.APIResponse("success get transactions by campaign", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, responseJSON)
}
