package handler

import (
	"campaignproject/campaign"
	"campaignproject/helper"
	"campaignproject/user"
	"net/http"
	"strconv"

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

func (h *campaignHandler) UserHaveCampaigns(c *gin.Context) {
	user := c.MustGet("user").(user.User)
	pagination := helper.Pagination{}
	err := c.Bind(&pagination)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("unproccesable entity", http.StatusUnprocessableEntity, "errors", data)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}
	campaigns, err := h.service.FindAllUserCampaign(user.ID, pagination)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("bad request", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	responseJSON := helper.APIResponse("get all user auth campaign", http.StatusOK, "succes", campaigns)
	c.JSON(http.StatusOK, responseJSON)
}

func (h *campaignHandler) FindAllCampaign(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	pagination := helper.Pagination{}
	err := c.Bind(&pagination)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("unproccesable entity", http.StatusUnprocessableEntity, "errors", data)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}
	campaigns, err := h.service.FindAllCampaign(userId, pagination)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("bad request", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	responseJSON := helper.APIResponse("success get all campaign", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, responseJSON)
}

func (h *campaignHandler) DetailBySlug(c *gin.Context) {
	slug := c.Param("slug")
	detail, err := h.service.DetailCampaignBySlug(slug)
	if err != nil {
		data := gin.H{"errors": err.Error()}
		responseJSON := helper.APIResponse("bad request", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	formatter := campaign.CreateDetailFormatter(detail)
	responseJSON := helper.APIResponse("get detail success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, responseJSON)
}
