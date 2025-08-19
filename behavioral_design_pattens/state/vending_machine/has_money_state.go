package vendingmachine

import (
	"fmt"
)

type HasMoneyState struct {
	vendingmachine *VendingMachine
}

func (s *HasMoneyState) AddItem(int) error {
	return fmt.Errorf("not allowed")
}

func (s *HasMoneyState) RequestItem() error {
	return fmt.Errorf("not allowed")
}

func (s *HasMoneyState) InsertMoney(int) error {
	return fmt.Errorf("not allowed")
}

func (s *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	s.vendingmachine.itemCount--

	if s.vendingmachine.itemCount == 0 {
		s.vendingmachine.SetState(&NoItemState{
			vendingmachine: s.vendingmachine,
		})
	} else {
		s.vendingmachine.SetState(&HasItemState{
			vendingmachine: s.vendingmachine,
		})
	}

	return nil
}
