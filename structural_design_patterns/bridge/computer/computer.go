package computer

import "lld/structural_design_patterns/bridge/printer"

type Computer interface {
	Print()
	SetPrinter(printer.Printer)
}
