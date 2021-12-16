package campaign

import (
	"strings"
	"time"
)

type CreateFormatter struct {
	ID               int       `json:"id"`
	UserId           int       `json:"user_id"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	ShortDescription string    `json:"short_description"`
	Perks            string    `json:"perks"`
	GoalAmount       int       `json:"goal_amount"`
	CurrentAmount    int       `json:"current_amount"`
	BeckerCount      int       `json:"becker_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type DetailCampaignFormatter struct {
	ID               int              `json:"id"`
	Slug             string           `json:"slug"`
	Image            string           `json:"images"`
	Perks            []string         `json:"perks"`
	ShortDescription string           `json:"short_description"`
	Description      string           `json:"description"`
	BackerCount      int              `json:"backer_count"`
	GoalAmount       int              `json:"goal_amount"`
	CurrentAmount    int              `json:"current_amount"`
	User             ProjectLoader    `json:"user"`
	CampaignImages   []ImageFormatter `json:"campaign_images"`
}

type ProjectLoader struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	AvatarFileName string `json:"avatar_file_name"`
}

type ImageFormatter struct {
	ID        int    `json:"id"`
	FileName  string `json:"file_name"`
	IsPrimary bool   `json:"is_primary"`
}
type CampaignFormatter struct {
	ID               int       `json:"id"`
	UserId           int       `json:"user_id"`
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

func CreateDetailFormatter(campaign Campaign) DetailCampaignFormatter {
	detailFormatter := DetailCampaignFormatter{}
	projectLoader := ProjectLoader{
		ID:             campaign.User.ID,
		Name:           campaign.User.Name,
		AvatarFileName: campaign.User.AvatarFileName,
	}

	perks := strings.Split(campaign.Perks, " ")

	imgFormatter := []ImageFormatter{}
	for _, image := range campaign.CampaignImages {
		isPrimary := false
		imageFormatter := ImageFormatter{}
		imageFormatter.ID = image.ID
		imageFormatter.FileName = image.FileName
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		imageFormatter.IsPrimary = isPrimary
		imgFormatter = append(imgFormatter, imageFormatter)
	}

	detailFormatter.ID = campaign.ID
	detailFormatter.Slug = campaign.Slug
	detailFormatter.ShortDescription = campaign.ShortDescription
	detailFormatter.Description = campaign.Description
	detailFormatter.BackerCount = campaign.BackerCount
	detailFormatter.GoalAmount = campaign.GoalAmount
	detailFormatter.CurrentAmount = campaign.CurrentAmount
	detailFormatter.User = projectLoader
	detailFormatter.CampaignImages = imgFormatter
	detailFormatter.Perks = perks

	detailFormatter.Image = ""
	if len(campaign.CampaignImages) != 0 {
		detailFormatter.Image = campaign.CampaignImages[0].FileName
	}
	return detailFormatter
}

func CreateFormat(campaign Campaign) CreateFormatter {
	formatter := CreateFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		Slug:             campaign.Slug,
		ShortDescription: campaign.ShortDescription,
		Perks:            campaign.Perks,
		BeckerCount:      campaign.BackerCount,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		CreatedAt:        campaign.CreatedAt,
		UpdatedAt:        campaign.UpdatedAt,
	}
	return formatter
}

func CreateCampaignFormatter(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserId = campaign.UserId
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
	return campaignFormatter
}

func CampaignsFormatter(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter := CreateCampaignFormatter(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}
