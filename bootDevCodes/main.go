package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?
func updateBalance(cus *customer, deposit transaction)error{
	if deposit.transactionType == transactionDeposit {
		cus.balance +=deposit.amount
	}else if deposit.transactionType == transactionWithdrawal{
		if cus.balance <= deposit.amount{
			return errors.New("insufficient funds")
		}
		cus.balance -=deposit.amount

	}else{
		return errors.New("unknown transaction type")
	}
	return nil
}