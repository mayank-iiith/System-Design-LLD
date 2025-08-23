package main

import (
	"fmt"
	"lld/behavioral_design_pattens/memento/caretaker"
	"lld/behavioral_design_pattens/memento/originators"
)

// Memento is a behavioral design pattern that lets you save and restore the previous state(snapshots) of an object without revealing the details of its implementation.

// The Memento doesn’t compromise the internal structure of the object it works with, as well as data kept inside the snapshots.

func main() {
	// ---- Bank Account Example ----
	account := originators.NewBankAccount(100)
	accountHistory := &caretaker.History{}

	accountHistory.Save(account)
	account.Deposit(15)
	accountHistory.Save(account)
	account.Withdraw(20)

	fmt.Println("Balance after ops:", account.Balance())

	// Undo
	accountHistory.Undo(account)
	fmt.Println("Balance after undo:", account.Balance())

	// Redo
	accountHistory.Redo(account)
	fmt.Println("Balance after redo:", account.Balance())

	fmt.Println()
	// ---- Document Editor Example ----
	doc := originators.NewDocumentEditor()
	docHistory := &caretaker.History{}

	docHistory.Save(doc)
	doc.Type("Hello, ")
	docHistory.Save(doc)
	doc.Type("World!")

	fmt.Println("Document:", doc.Content())

	docHistory.Undo(doc)
	fmt.Println("After Undo:", doc.Content())

	docHistory.Redo(doc)
	fmt.Println("After Redo:", doc.Content())
}
