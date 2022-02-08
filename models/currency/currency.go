package currency

import (
	"github.com/leekchan/accounting"
)

func (currency *Currency) ToRand() string {
	ac := accounting.Accounting{Symbol: "R", Precision: 2}

	return ac.FormatMoney(currency.Amount)
}

type Currency struct {
	Amount int32
}
