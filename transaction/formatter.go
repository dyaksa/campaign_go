package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	var campaignTransactionsFormatter []CampaignTransactionFormatter
	for _, transaction := range transactions {
		campaignTransactionFormatter := FormatCampaignTransaction(transaction)
		campaignTransactionsFormatter = append(campaignTransactionsFormatter, campaignTransactionFormatter)
	}
	return campaignTransactionsFormatter
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Amount    int               `json:"amount"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatterUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormat := CampaignFormatter{}
	campaignFormat.Name = "test"
	campaignFormat.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormat.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}
	formatter.Campaign = campaignFormat
	return formatter
}

func FormatterUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	var transactionsFormatter []UserTransactionFormatter
	for _, transaction := range transactions {
		formatter := FormatterUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}
