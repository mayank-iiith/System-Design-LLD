package payment

import "fmt"

type Ledger struct {
}

func NewLedger() *Ledger {
	return &Ledger{}
}

func (l *Ledger) MakeEntry(accountName, txnType string, txnAmount int) {
	fmt.Printf("Ledger entry created for accountName %q with txnType %q for amount %d\n", accountName, txnType, txnAmount)
}
