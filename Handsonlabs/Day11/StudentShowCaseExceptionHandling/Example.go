/*
Ecommerce Example For Cart
Bounderies

	Minium 1 items should be available in cart
	able add items or  delete items
	able to increae and decrese items

Checkout Pyload
Total based on cart items
Tax need to applied
Discount coupon need to applied

//Method of Payment
PaymentMode

	{
		Type String
	}

Card
UPI
NEFT
ETC

interface to provide common and specific functionaliy to each strucutre

Exception Handling
validate the different structs
Validate the Cart
Validate Payment
Panic for System Failure

Quesiton ?
How we can define this System?
What tools from Go lang you will use here and why ?
*/
package main

import (
	"errors"
	"fmt"
)

//Cart
// Products
//Qua

//Product
/**
Price
ComapnyName
Name
Discount

**/
//Payment

func main() {
	fmt.Println("Cart System")
	paymentMode := paymentMode{
		TransactionId: "dsfdsadsasa",
		Type:          "No Payment Type Seclected",
	}
	// card := card{
	// 	PaymentMode: paymentMode,
	// 	CardNumber:  "1234-5678-9012-3456",
	// }
	// upiMethod := upi{
	// 	PaymentMode: paymentMode,
	// 	UpiNumber:   "user@upi",

	// }
	netBanking := netBanking{
		PaymentMode:   paymentMode,
		AccountNumber: "8383838383838",
	}
	// if card.getType() == "Card" {
	// 	err := card.validatePayment()
	// 	fmt.Println("Card Payment Validation Error:", err)
	// }

	validateByType(netBanking)

	// err := card.validatePayment()
	// if err != nil {
	// 	fmt.Println("Error validating card:", err)
	// }
	// err = upiMethod.validatePayment()
	// if err != nil {
	// 	fmt.Println("Error validating UPI:", err)
	// }
	// err = netBanking.validatePayment()
	// if err != nil {
	// 	fmt.Println("Error validating Net Banking:", err)
	// }
}

type product struct {
	ID          int
	Price       float32
	CompanyName string
	Discount    int
}
type cart struct {
	Products []product
}
type paymentMode struct {
	TransactionId string
	Type          string
}

type card struct {
	PaymentMode paymentMode
	Pin         string
	CardNumber  string
	Expirydate  string
}

type netBanking struct {
	PaymentMode   paymentMode
	Pin           string
	AccountNumber string
	IFCCode       string
}
type upi struct {
	PaymentMode paymentMode
	UpiNumber   string
}

func (paymentMode paymentMode) validatePayment() error {
	if paymentMode.TransactionId == "" {
		return errors.New("Transaction Id is required")
	} else if paymentMode.Type == "" {
		return errors.New("Type  is required")
	}
	return nil
}

func (paymentMode card) validatePayment() error {
	err := paymentMode.PaymentMode.validatePayment()
	if err != nil {
		return fmt.Errorf("Error while validating Card: %w", err)
	} else if paymentMode.CardNumber == "" {
		return errors.New("Card Number is required")
	}
	return nil
}

func (upi upi) validatePayment() error {
	err := upi.PaymentMode.validatePayment()
	if err != nil {
		return fmt.Errorf("Error while validating UPI: %w", err)
	} else if upi.UpiNumber == "" {
		return errors.New("UpI Number is required")
	}
	return nil
}
func (netBankingMode netBanking) validatePayment() error {
	err := netBankingMode.PaymentMode.validatePayment()
	if err != nil {
		return fmt.Errorf("Error while validating Net Banking: %w", err)
	} else if netBankingMode.AccountNumber == "" {
		return errors.New("Account Number is required")
	}
	return nil
}

type validatePaymentMethod interface {
	validateMethod() error
}
type typeChecker interface {
	getType() string
}

func (upiMethod upi) validateMethod() error {
	return upiMethod.validatePayment()
}

func (cardMethod card) validateMethod() error {
	return cardMethod.validatePayment()
}
func (netBankingMethod netBanking) validateMethod() error {
	return netBankingMethod.validatePayment()
}

func (PaymentMode card) getType() string {
	return "Card"
}

func (PaymentMode upi) getType() string {
	return "UPI"
}
func (PaymentMode netBanking) getType() string {
	return "NetBanking"
}

// func validateByType(TypeChecker typeChecker, validator validatePaymentMethod) error{
// 	type := typeChecker.getType()
// 	switch type {
// 	case "Card":
// 		return validator.validateMethod()
// 	case "UPI":
// 		return validator.validateMethod()
// 	case "NetBanking":
// 		return validator.validateMethod()
// 	default:
// 		return errors.New("Unknown Payment Type")
// 	}
// }

func validateByType(validator validatePaymentMethod) error {
	err := validator.validateMethod()
	if err != nil {
		fmt.Printf("Error in ValidationBy Type: %w", err)
	}
	return err

}
