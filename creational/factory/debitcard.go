package factory

import "fmt"

type DebitCardPM struct{}

func (c *DebitCardPM) Pay(amt float32) string {
	return fmt.Sprintf("%0.2f paid by debit card", amt)
}
