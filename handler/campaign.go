package handler

import (
	"campaignproject/campaign"
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
	var input campaign.CampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	campaign, err := h.service.InputInsertCampaign(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, campaign)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
