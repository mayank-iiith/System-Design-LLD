package main

import (
	"fmt"
	"lld/structural_design_patterns/bridge/computer"
	"lld/structural_design_patterns/bridge/printer"
)

// Bridge is a structural design pattern that lets you split a large class or a set of closely related classes into two separate hierarchies—abstraction and implementation—which can be developed independently of each other.

// One of these hierarchies (often called the Abstraction) will get a reference to an object of the second hierarchy (Implementation). The abstraction will be able to delegate some (sometimes, most) of its calls to the implementations object. Since all implementations will have a common interface, they’d be interchangeable inside the abstraction.

func main() {
	hpPrinter := &printer.Hp{}
	epsonPrinter := &printer.Epson{}

	macComputer := &computer.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &computer.Windows{}

	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()
}
