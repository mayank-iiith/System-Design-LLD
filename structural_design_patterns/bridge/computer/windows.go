package computer

import (
	"fmt"
	"lld/structural_design_patterns/bridge/printer"
)

type Windows struct {
	printer printer.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p printer.Printer) {
	w.printer = p
}
