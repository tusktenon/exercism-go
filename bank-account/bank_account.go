package account

import "sync"

type Account struct {
	mu      sync.Mutex // guards account
	balance int64
	closed  bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance, !a.closed
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return 0, false
	}
	newBalance := a.balance + amount
	if newBalance < 0 {
		return a.balance, false
	}
	a.balance = newBalance
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return 0, false
	}
	closingBalance := a.balance
	a.balance = 0
	a.closed = true
	return closingBalance, true
}
