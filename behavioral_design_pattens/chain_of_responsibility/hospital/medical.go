package hospital

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}

func (m *Medical) Execute(p *Patient) {
	if p.DoctorCheckupDone {
		fmt.Printf("Medicine already given to patient %q \n", p.Name)
		m.next.Execute(p)
		return
	}
	fmt.Printf("Medical giving medicine to patient %q \n", p.Name)
	p.RegistrationDone = true
	m.next.Execute(p)
}
