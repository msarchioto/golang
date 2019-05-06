package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

// Wallet Bitcoin struct
type Wallet struct {
	balance Bitcoin
}

// String() Stringer interface implementation for Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit Wallet function
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance Wallet function
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw Wallet function
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func main() {
	fmt.Println("Bitcoin Wallet Pointers Test")
}
