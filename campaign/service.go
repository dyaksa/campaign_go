package campaign

import (
	"campaignproject/helper"
	"campaignproject/user"
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	InputInsertCampaign(input CampaignInput) (Campaign, error)
	FindAllUserCampaign(ID int, paginate helper.Pagination) (*helper.Pagination, error)
	FindAllCampaign(UserId int, paginate helper.Pagination) (*helper.Pagination, error)
	DetailCampaignBySlug(slug string) (Campaign, error)
	UpdateCampaign(InputID user.User, campaignID DetailCampaignInputId, input CampaignInput) (Campaign, error)
	UploadCampaignImages(input CampaignImagesInput, FileName string) (CampaignImages, error)
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
	campaign.User = input.User
	campaign.Description = input.Description
	campaign.ShortDescription = input.ShortDescription
	campaign.Perks = input.Perks

	titleSlug := fmt.Sprintf("%s-%s", input.Name, time)
	campaign.Slug = slug.Make(titleSlug)

	campaign.BackerCount = 0
	campaign.CurrentAmount = 0
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

func (s *service) FindAllCampaign(UserId int, paginate helper.Pagination) (*helper.Pagination, error) {
	if UserId != 0 {
		campaigns, err := s.repository.FindByUserId(UserId, paginate)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll(paginate)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) DetailCampaignBySlug(slug string) (Campaign, error) {
	campaign, err := s.repository.FindBySlug(slug)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) UpdateCampaign(InputID user.User, campaignID DetailCampaignInputId, input CampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(campaignID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != InputID.ID {
		return campaign, errors.New("user cannot access this campaign")
	}

	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.BackerCount = 0
	campaign.GoalAmount = input.GoalAmount
	campaign.User = InputID
	saved, err := s.repository.Updated(campaign)
	if err != nil {
		return saved, err
	}
	return saved, nil

}

func (s *service) UploadCampaignImages(input CampaignImagesInput, FileName string) (CampaignImages, error) {
	campaignImages := CampaignImages{}
	isPrimary := 0
	campaign, err := s.repository.FindById(input.CampaignID)

	if err != nil {
		return campaignImages, err
	}

	if campaign.UserId != input.User.ID {
		return campaignImages, errors.New("user cannot upload campaign images")
	}

	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllCampaignImegesIsPrimaryFalse(input.CampaignID)
		if err != nil {
			return campaignImages, err
		}
	}
	campaignImages.IsPrimary = isPrimary
	campaignImages.CampaignId = input.CampaignID
	campaignImages.FileName = FileName
	saved, err := s.repository.SaveCampaignImages(campaignImages)
	if err != nil {
		return campaignImages, err
	}
	return saved, nil
}
