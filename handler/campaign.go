package handler

import (
	"campaignproject/campaign"
	"campaignproject/helper"
	"campaignproject/user"
	"fmt"
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
	currentUser := c.MustGet("user").(user.User)
	var input campaign.CampaignInput
	input.User = currentUser
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
	param := campaign.DetailCampaignInput{}
	err := c.ShouldBindUri(&param)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		responseJSON := helper.APIResponse("bad request error", http.StatusBadRequest, "errors", errors)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}

	detail, err := h.service.DetailCampaignBySlug(param.Slug)
	if detail.ID == 0 {
		responseJSON := helper.APIResponse("not found", http.StatusNotFound, "errors", nil)
		c.JSON(http.StatusNotFound, responseJSON)
		return
	}

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

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	currentUser := c.MustGet("user").(user.User)
	campaignID := campaign.DetailCampaignInputId{}
	input := campaign.CampaignInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		responseJSON := helper.APIResponse("updated failed", http.StatusBadRequest, "errors", errors)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}

	err = c.ShouldBindUri(&campaignID)
	if err != nil {
		responseJSON := helper.APIResponse("updated failed", http.StatusBadRequest, "errors", nil)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}

	newCampaign, err := h.service.UpdateCampaign(currentUser, campaignID, input)
	if err != nil {
		responseJSON := helper.APIResponse("updated campaign failed", http.StatusBadRequest, "errors", err.Error())
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	formatter := campaign.CreateFormat(newCampaign)
	responseJSON := helper.APIResponse("success updated campaign", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, responseJSON)
}

func (h *campaignHandler) UploadCampaignImages(c *gin.Context) {
	user := c.MustGet("user").(user.User)
	input := campaign.CampaignImagesInput{}
	input.User = user
	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		responseJSON := helper.APIResponse("uploaded image faile", http.StatusUnprocessableEntity, "errors", data)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}

	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatterErroValidation(err)
		responseJSON := helper.APIResponse("upload image failed", http.StatusUnprocessableEntity, "errors", errors)
		c.JSON(http.StatusUnprocessableEntity, responseJSON)
		return
	}
	path := fmt.Sprintf("images/%d-%s", user.ID, file.Filename)
	c.SaveUploadedFile(file, path)
	_, err = h.service.UploadCampaignImages(input, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		responseJSON := helper.APIResponse(err.Error(), http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, responseJSON)
		return
	}
	data := gin.H{"is_uploaded": true}
	responseJSON := helper.APIResponse("uploaded campaign images success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, responseJSON)
}
