package bankaccount

import (
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func New() *BankAccount {
	return &BankAccount{
		balance: 0,
		mu:      sync.Mutex{},
	}
}

func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amount
}

func (a *BankAccount) GetBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}
