package factory

import "fmt"

type CashPM struct{}

func (c *CashPM) Pay(amt float32) string {
	return fmt.Sprintf("%0.2f paid by cash", amt)
}
