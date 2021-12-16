package campaign

import "campaignproject/user"

type CampaignInput struct {
	User             user.User
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	BackerCount      int    `json:"backer_count" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	CurrentAmount    int    `json:"current_amount"`
}

type DetailCampaignInput struct {
	Slug string `uri:"slug" binding:"required"`
}

type DetailCampaignInputId struct {
	ID int `uri:"id" binding:"required"`
}
