package payment

import "fmt"

type Notification struct {
}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) SendWalletCreditNotification(accountName string, amount int) {
	fmt.Println("Account %q credited with amount %d", accountName, amount)
}

func (n *Notification) SendWalletDebitNotification(accountName string, amount int) {
	fmt.Println("Account %q debited with amount %d", accountName, amount)
}
