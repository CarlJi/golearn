package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(count float32) string
}

const (
	Card  = 10.0
	Debit = 30.0
)

type CardMethod struct {
}

func (c *CardMethod) Pay(count float32) string {
	return fmt.Sprintf("using %v", CardMethod{})
}

type DebitMethod struct {
}

func (d *DebitMethod) Pay(count float32) string {
	return fmt.Sprintf("using %v", DebitMethod{})
}

func GetPaymentMethod(count int) (PaymentMethod, error) {
	switch count {
	case Card:
		return new(CardMethod), nil
	case Debit:
		return new(DebitMethod), nil

	default:
		return nil, errors.New("unkown payment type")
	}
}
