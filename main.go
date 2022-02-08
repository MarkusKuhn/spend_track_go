package main

import (
	"fmt"
	"spend_track_go/config/server"
	"spend_track_go/models/currency"
	"spend_track_go/models/debit"
)

func main() {
	d := debit.Debit{Amount: currency.Currency{Amount: 12312}, Description: "Test"}
	fmt.Println("Amount: ")
	fmt.Println(d.Amount.ToRand())
	server.Start()
}
