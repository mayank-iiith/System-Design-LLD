package payment

import "fmt"

type Account struct {
	name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) VerifyAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name %q is invalid", accountName)
	}
	fmt.Printf("Account Verified")
	return nil
}
