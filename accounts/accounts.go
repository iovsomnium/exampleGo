package accounts

import "errors"

//Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

var errNoMoney = errors.New("Can't withdraw")

// Withdraw - amount on your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil //python null
}

// Balance return balance
func (a Account) Balance() int {
	return a.balance
}
