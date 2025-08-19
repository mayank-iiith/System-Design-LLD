package vendingmachine

import (
	"fmt"
)

// NoItemState - In this state we can either add more item, or request the added item
type HasItemState struct {
	vendingmachine *VendingMachine
}

func (s *HasItemState) AddItem(count int) error {
	fmt.Printf("Adding %d more items\n", count)
	s.vendingmachine.IncrementItemCount(count)
	return nil
}

func (s *HasItemState) RequestItem() error {
	if s.vendingmachine.itemCount == 0 {
		s.vendingmachine.SetState(&NoItemState{
			vendingmachine: s.vendingmachine,
		})
	}
	fmt.Println("Item Requested")
	s.vendingmachine.SetState(&ItemRequestedState{
		vendingmachine: s.vendingmachine,
	})
	return nil
}

func (s *HasItemState) InsertMoney(int) error {
	return fmt.Errorf("not allowed")
}

func (s *HasItemState) DispenseItem() error {
	return fmt.Errorf("not allowed")
}
