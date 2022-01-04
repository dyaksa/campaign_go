package transaction

import "campaignproject/user"

type GetCampaignID struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type InputTransaction struct {
	UserID     int `binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	Amount     int `json:"amount" binding:"required"`
	Status     string
	Email      string
	Name       string
}

type TransactionProcessInput struct {
	OrderID     string `query:"order_id"`
	PaymentType string `query:"payment_type"`
	Amount      int    `query:"gross_amount"`
	Status      string `query:"transaction_status"`
	FraudStatus string `query:"fraud_status"`
}
