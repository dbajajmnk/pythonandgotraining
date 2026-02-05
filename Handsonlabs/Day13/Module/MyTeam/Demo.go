package main

import (
	"fmt"
    "github.com/durango/go-credit-card"
)

func main() {

	card := creditcard.Card{Number: "4242424242424242", Cvv: "11111", Month: "02", Year: "2016"}

	// Retrieve the card's method (which credit card company this card belongs to)
	err := card.Method() // card.Company({Short: "visa", Long: "Visa"})

	// Display last four digits
	lastFour, err := card.LastFour() // 4242

	// Validate the card's number (without capturing)
	err := card.Validate() // will return an error due to not allowing test cards

	err := card.Validate(true) // this will work though
	fmt.Println("Error", err)

}
