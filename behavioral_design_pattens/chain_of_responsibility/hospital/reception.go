package hospital

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Printf("Patient %q registration is already done \n", p.Name)
		r.next.Execute(p)
		return
	}
	fmt.Printf("Reception registering patient %q \n", p.Name)
	p.RegistrationDone = true
	r.next.Execute(p)
}
