package pointersnerrors

import (
	"errors"
	"fmt"
)

var ErrInsuficientFounds = errors.New("cannot withdraw, insufficient founds")

type Money int

type Wallet struct {
	balance Money
}

func (wallet *Wallet) Deposit(amount Money) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Money {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Money) error {
	if amount > wallet.balance {
		return ErrInsuficientFounds
	}

	wallet.balance -= amount
	return nil
}

type Stringer interface {
	String() string
}

func (m Money) String() string {
	return fmt.Sprintf("%d BRL", m)
}
