package transaction

import (
	"campaignproject/campaign"
	"campaignproject/helper"
	"errors"
)

type Service interface {
	GetTransacationsByCampaignID(input GetCampaignID, paginate helper.Pagination) (*helper.Pagination, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository: repository, campaignRepository: campaignRepository}
}

func (s *service) GetTransacationsByCampaignID(input GetCampaignID, paginate helper.Pagination) (*helper.Pagination, error) {
	transactions, err := s.repository.GetByCampaignID(input, paginate)
	if err != nil {
		return transactions, err
	}

	campaign, err := s.campaignRepository.FindById(input.ID)
	if err != nil {
		return transactions, err
	}

	if campaign.UserId != input.User.ID {
		return transactions, errors.New("user has not access campaign id")
	}

	return transactions, err
}
