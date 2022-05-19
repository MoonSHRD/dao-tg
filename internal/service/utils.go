package service

import (
	"math/big"
	"time"

	"github.com/shopspring/decimal"
)

// Kinda zeroed time - 1 Jan 1970
var motherBirth = time.UnixMilli(0)

// Returns tg.Recipient
func newRecipient(value string) pocketRecipient {
	return pocketRecipient{value}
}

// tg.Recipient compat
type pocketRecipient struct {
	recipient string
}

func (r pocketRecipient) Recipient() string {
	return r.recipient
}

const defaultEtherDecimals = 18

// toDecimal wei to decimals
func toDecimal(value *big.Int, decimals int) decimal.Decimal {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	return num.Div(mul)
}

func strToDecimal(value string, decimals int) decimal.Decimal {
	val := new(big.Int)
	val.SetString(value, 10)
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(val.String())
	return num.Div(mul)
}
