package transaction

import (
	"campaignproject/campaign"
	"campaignproject/helper"
	"campaignproject/payment"
	"errors"
	"time"
)

type Service interface {
	GetTransacationsByCampaignID(input GetCampaignID, paginate helper.Pagination) (*helper.Pagination, error)
	GetTransactionsByUserID(campaignID int, paginate helper.Pagination) (*helper.Pagination, error)
	SaveTransaction(input InputTransaction) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository: repository, campaignRepository: campaignRepository, paymentService: paymentService}
}

func (s *service) GetTransacationsByCampaignID(input GetCampaignID, paginate helper.Pagination) (*helper.Pagination, error) {
	transactions, err := s.repository.GetByCampaignID(input.ID, paginate)
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

func (s *service) GetTransactionsByUserID(UserID int, paginate helper.Pagination) (*helper.Pagination, error) {
	transactions, err := s.repository.GetByUserId(UserID, paginate)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) SaveTransaction(input InputTransaction) (Transaction, error) {
	transaction := Transaction{}
	existCampaign, err := s.campaignRepository.FindById(input.CampaignID)
	if err != nil {
		return transaction, err
	}
	if existCampaign.ID == 0 {
		return transaction, errors.New("campaign id not found")
	}
	transaction.UserID = input.UserID
	transaction.Status = "pending"
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.Code = "ORDER-" + time.Now().Format("20060102150405")
	transaction, err = s.repository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	paymentInput := payment.Transaction{
		Name:   input.Name,
		Email:  input.Email,
		Code:   transaction.Code,
		Amount: input.Amount,
	}
	paymentURL, err := s.paymentService.GetToken(paymentInput)
	if err != nil {
		return transaction, err
	}
	transaction.PaymentURL = paymentURL
	transaction, err = s.repository.Update(transaction)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
