package debit

import (
	"github.com/jinzhu/gorm"
	money "spend_track_go/models/currency"
)

type Debit struct {
	gorm.Model

	Amount      money.Currency
	Description string
}
