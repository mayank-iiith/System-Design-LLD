package payment

import "fmt"

type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Amount credited from wallet balance")
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	w.balance -= amount
	fmt.Println("Amount deducted from wallet balance")
	return nil
}
