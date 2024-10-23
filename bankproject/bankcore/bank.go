package bank

import (
	"errors"
	"fmt"
)

// 顾客 ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// 用户 ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

type Bank interface {
	Statement() string
}

// 存储 ...
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be less than the account's balance")
	}

	a.Balance -= amount
	return nil
}

func Statement(b Bank) string {
	return b.Statement()
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(o *Account, amount float64) error {
	if a.Balance < amount {
		return errors.New("the balance to transfer should be greater than amount")
	}
	a.Withdraw(amount)
	o.Deposit(amount)
	return nil
}
