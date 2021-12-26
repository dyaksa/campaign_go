package transaction

type Service interface {
	GetTransacationsByCampaignID(campaignID int) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetTransacationsByCampaignID(campaignID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByCampaignID(campaignID)
	if err != nil {
		return transactions, err
	}
	return transactions, err
}
