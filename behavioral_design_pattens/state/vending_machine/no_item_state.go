package vendingmachine

import (
	"fmt"
)

// NoItemState - In this state we can only add item
type NoItemState struct {
	vendingmachine *VendingMachine
}

func (s *NoItemState) AddItem(count int) error {
	fmt.Printf("Adding %d items\n", count)
	s.vendingmachine.IncrementItemCount(count)
	s.vendingmachine.SetState(&HasItemState{
		vendingmachine: s.vendingmachine,
	})
	return nil
}

func (s *NoItemState) RequestItem() error {
	return fmt.Errorf("not allowed")
}

func (s *NoItemState) InsertMoney(int) error {
	return fmt.Errorf("not allowed")
}

func (s *NoItemState) DispenseItem() error {
	return fmt.Errorf("not allowed")
}
