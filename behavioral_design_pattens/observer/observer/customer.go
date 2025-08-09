package observer

import "fmt"

type Customer struct {
	id string
}

func NewCustomer(id string) *Customer {
	return &Customer{
		id: id,
	}
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %q for item %q\n", c.id, itemName)
}

func (c *Customer) GetId() string {
	return c.id
}
