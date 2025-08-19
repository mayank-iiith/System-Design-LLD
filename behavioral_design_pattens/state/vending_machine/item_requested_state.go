package vendingmachine

import (
	"fmt"
)

type ItemRequestedState struct {
	vendingmachine *VendingMachine
}

func (s *ItemRequestedState) AddItem(int) error {
	return fmt.Errorf("not allowed")
}

func (s *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("not allowed")
}

func (s *ItemRequestedState) InsertMoney(money int) error {
	if s.vendingmachine.itemPrice < money {
		return fmt.Errorf("inserted money is less, please insert %d", s.vendingmachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	s.vendingmachine.SetState(&HasMoneyState{
		vendingmachine: s.vendingmachine,
	})
	return nil
}

func (s *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("not allowed")
}
