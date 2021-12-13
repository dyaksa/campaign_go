package campaign

import (
	"campaignproject/helper"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	InputInsertCampaign(input CampaignInput) (Campaign, error)
	FindAllUserCampaign(ID int, paginate helper.Pagination) (*helper.Pagination, error)
	FindAllCampaign(paginate helper.Pagination) (*helper.Pagination, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) InputInsertCampaign(input CampaignInput) (Campaign, error) {
	time := time.Now().Format("20060102150405")
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.UserId = input.UserId
	campaign.Description = input.Description
	campaign.ShortDescription = input.ShortDescription
	campaign.Perks = input.Perks

	titleSlug := fmt.Sprintf("%s-%s", input.Name, time)
	campaign.Slug = slug.Make(titleSlug)

	campaign.BackerCount = input.BackerCount
	campaign.CurrentAmount = input.CurrentAmount
	campaign.GoalAmount = input.GoalAmount
	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return campaign, err
	}
	return newCampaign, nil
}

func (s *service) FindAllUserCampaign(ID int, paginate helper.Pagination) (*helper.Pagination, error) {
	campaigns, err := s.repository.FindByUserId(ID, paginate)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) FindAllCampaign(paginate helper.Pagination) (*helper.Pagination, error) {
	campaigns, err := s.repository.FindAll(paginate)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
