package payment

import "fmt"

type WalletFacade struct {
	account      *Account
	securityCode *SecurityCode
	wallet       *Wallet
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountName string, securityCode int) *WalletFacade {
	return &WalletFacade{
		account:      NewAccount(accountName),
		securityCode: NewSecurityCode(securityCode),
		wallet:       NewWallet(),
		notification: NewNotification(),
		ledger:       NewLedger(),
	}
}

func (w *WalletFacade) AddMonetToWallet(accountName string, securityCode, amount int) error {
	fmt.Println("Starting adding money to wallet")
	if err := w.account.VerifyAccount(accountName); err != nil {
		return err
	}
	if err := w.securityCode.VerifySecurityCode(securityCode); err != nil {
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification(accountName, amount)
	w.ledger.MakeEntry(accountName, "CREDIT", amount)
	return nil
}

func (w *WalletFacade) DeductMonetFromWallet(accountName string, securityCode, amount int) error {
	fmt.Println("Starting deducting money from wallet")
	if err := w.account.VerifyAccount(accountName); err != nil {
		return err
	}
	if err := w.securityCode.VerifySecurityCode(securityCode); err != nil {
		return err
	}
	err := w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.SendWalletDebitNotification(accountName, amount)
	w.ledger.MakeEntry(accountName, "DEBIT", amount)
	return nil
}
