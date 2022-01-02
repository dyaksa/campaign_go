package payment

import (
	"os"

	midtrans "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Service interface {
	GetToken(transaction Transaction) (string, error)
}

type service struct {
}

func NewPaymantService() *service {
	return &service{}
}

func (s *service) GetToken(transaction Transaction) (string, error) {
	var midSnap snap.Client
	midSnap.New(os.Getenv("MIDTRANS_ACCESS_KEY"), midtrans.Sandbox)
	params := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Code,
			GrossAmt: int64(transaction.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transaction.Name,
			Email: transaction.Email,
		},
	}
	snapResp, err := midSnap.CreateTransaction(params)
	if err != nil {
		return "", err.RawError
	}
	return snapResp.RedirectURL, nil
}
