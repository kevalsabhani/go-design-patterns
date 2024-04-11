package factory

import (
	"errors"
)

const (
	Cash      = 1
	DebitCard = 2
)

type PaymentMethod interface {
	Pay(amt float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New("Payment method not recognized")
	}
}
