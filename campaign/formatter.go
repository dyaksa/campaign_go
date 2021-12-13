package campaign

import (
	"time"
)

type CreateFormatter struct {
	ID               int       `json:"id"`
	UserId           int       `json:"user_id"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	ShortDescription string    `json:"short_description"`
	Perks            string    `json:"perks"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ListFormatter struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	Image            string    `json:"images"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	BackerCount      int       `json:"backer_count"`
	GoalAmount       int       `json:"goal_amount"`
	CurrentAmount    int       `json:"current_amount"`
	CreatedAt        time.Time `json:"created_at"`
}

func CreateFormat(campaign Campaign) CreateFormatter {
	formatter := CreateFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		Slug:             campaign.Slug,
		ShortDescription: campaign.ShortDescription,
		Perks:            campaign.Perks,
		CreatedAt:        campaign.CreatedAt,
		UpdatedAt:        campaign.UpdatedAt,
	}
	return formatter
}

func CreateListFormatter(campaigns []Campaign) []ListFormatter {
	formatter := []ListFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter := ListFormatter{}
		campaignFormatter.ID = campaign.ID
		campaignFormatter.Name = campaign.Name
		campaignFormatter.Slug = campaign.Slug
		campaignFormatter.ShortDescription = campaign.ShortDescription
		campaignFormatter.Description = campaign.Description
		campaignFormatter.BackerCount = campaign.BackerCount
		campaignFormatter.Image = ""
		if len(campaign.CampaignImages) > 0 {
			campaignFormatter.Image = campaign.CampaignImages[0].FileName
		}
		campaignFormatter.GoalAmount = campaign.GoalAmount
		campaignFormatter.CurrentAmount = campaign.CurrentAmount
		campaignFormatter.CreatedAt = campaign.CreatedAt
		formatter = append(formatter, campaignFormatter)
	}
	return formatter
}
