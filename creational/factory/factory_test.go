package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	assert.Nil(t, err, "Payment method not recognized")
	assert.NotNil(t, payment)
	msg := payment.Pay(10.0)
	assert.Contains(t, msg, "paid by cash")
}

func TestPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	assert.Nil(t, err, "Payment method not recognized")
	assert.NotNil(t, payment)
	msg := payment.Pay(10.0)
	assert.Contains(t, msg, "paid by debit card")
}

func TestPaymentMethodNotExistent(t *testing.T) {
	_, err := GetPaymentMethod(0)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Payment method not recognized")
}
