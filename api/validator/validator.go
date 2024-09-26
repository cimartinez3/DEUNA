package validator

import (
	"errors"
	"strings"
	"unicode"

	"github.com/cimartinez3/DEUNA"
	"github.com/cimartinez3/DEUNA/api/types"
)

// IValidator interface that defines methods for validations.
type IValidator interface {
	ValidateTransaction(req types.ChargeRequest) error
}

// Charge struct that will implement IValidator.
type Charge struct{}

// NewChargeValidator creates a new Charge struct validator.
func NewChargeValidator() IValidator {
	return &Charge{}
}

// ValidateTransaction makes fields validations for charge request.
func (c *Charge) ValidateTransaction(req types.ChargeRequest) error {
	if strings.EqualFold(DEUNA.MockBadCard, req.Card.Number) {
		return errors.New("this bank does not support this card")
	}

	if !onlyNumbers(req.Card.Number) {
		return errors.New("wrong card data")

	}

	if req.Amount <= 0 {
		return errors.New("amount cant be less than 0")
	}

	return nil
}

// onlyNumbers used to validate if card contains words.
func onlyNumbers(card string) bool {
	for _, char := range card {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
