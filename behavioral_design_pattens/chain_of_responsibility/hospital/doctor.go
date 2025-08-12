package hospital

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckupDone {
		fmt.Printf("Doctor checkup is already done for patient %q \n", p.Name)
		d.next.Execute(p)
		return
	}
	fmt.Printf("Doctor checking patient %q \n", p.Name)
	p.RegistrationDone = true
	d.next.Execute(p)
}
