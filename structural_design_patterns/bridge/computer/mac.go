package computer

import (
	"fmt"
	"lld/structural_design_patterns/bridge/printer"
)

type Mac struct {
	printer printer.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p printer.Printer) {
	m.printer = p
}
