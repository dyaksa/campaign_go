package handler

import (
	"campaignproject/campaign"
	"campaignproject/helper"
	"campaignproject/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service: service}
}

func (h *campaignHandler) InputInsertCampaign(c *gin.Context) {
	user := c.MustGet("user").(user.User)
	var input campaign.CampaignInput
	input.UserId = user.ID
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		data := gin.H{"errors": errors}
		responseJSON := helper.APIResponse("bad request", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}

	newCampaign, err := h.service.InputInsertCampaign(input)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("bad request", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	formatter := campaign.CreateFormat(newCampaign)
	responseJSON := helper.APIResponse("success create campaign", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusOK, responseJSON)
}
