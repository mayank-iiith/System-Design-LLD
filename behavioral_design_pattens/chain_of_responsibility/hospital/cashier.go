package hospital

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}

func (c *Cashier) Execute(p *Patient) {
	if p.DoctorCheckupDone {
		fmt.Printf("Payment Done for patient %q \n", p.Name)
		return
	}
	fmt.Printf("Cashier getting money from patient %q \n", p.Name)
}
