package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("widthrawl amount is higher than the wallet's balance")

type Stringer interface {
	String() string
}

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Widthraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
