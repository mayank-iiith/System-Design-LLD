package vendingmachine

type VendingMachine struct {
	currentState State

	itemCount int
	itemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	if itemCount == 0 {
		v.SetState(&NoItemState{
			vendingmachine: v,
		})
	} else {
		v.SetState(&HasItemState{
			vendingmachine: v,
		})
	}

	return v
}

func (v *VendingMachine) SetState(state State) {
	v.currentState = state
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}
func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}
func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}
func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) IncrementItemCount(count int) {
	v.itemCount += count
}
