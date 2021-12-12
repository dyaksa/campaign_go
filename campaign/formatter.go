package campaign

import "time"

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
