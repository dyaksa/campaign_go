package campaign

import (
	"github.com/gosimple/slug"
)

type Service interface {
	InputInsertCampaign(input CampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) InputInsertCampaign(input CampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.Description = input.Description
	campaign.ShortDescription = input.ShortDescription
	campaign.Perks = input.Perks
	campaign.Slug = slug.Make(input.Name)
	campaign.BackerCount = input.BackerCount
	campaign.CurrentAmount = input.CurrentAmount
	campaign.GoalAmount = input.GoalAmount
	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return campaign, err
	}
	return newCampaign, nil
}
