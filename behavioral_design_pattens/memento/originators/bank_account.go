package originators

import (
	"lld/behavioral_design_pattens/memento/types"
)

// BankAccount is a concrete originator
type BankAccount struct {
	balance int
}

// BankAccountMemento is the snapshot of BankAccount
type BankAccountMemento struct {
	Balance int
}

func NewBankAccount(initalBalance int) *BankAccount {
	return &BankAccount{
		balance: initalBalance,
	}
}

func (a *BankAccount) Deposit(amount int) {
	a.balance += amount
}

func (a *BankAccount) Withdraw(amount int) {
	a.balance -= amount
}

func (a *BankAccount) Balance() int {
	return a.balance
}

// Implement Originator interface
var _ types.Originator = (*BankAccount)(nil)

// TakeSnapshot implements types.Originator.
func (a *BankAccount) TakeSnapshot() types.Memento {
	return BankAccountMemento{
		Balance: a.balance,
	}
}

// Restore implements types.Originator.
func (a *BankAccount) Restore(m types.Memento) {
	if snap, ok := m.(BankAccountMemento); ok {
		a.balance = snap.Balance
	}
}
