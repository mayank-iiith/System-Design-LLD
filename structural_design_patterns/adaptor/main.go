package main

import "fmt"

// Adapter is a structural design pattern that allows objects with incompatible interfaces/objects to collaborate.
// It is also known as "Wrapper"
// The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.

func main() {
	client := &Client{}

	mac := &Mac{}
	client.InsertLightningPortIntoComputer(mac)
	fmt.Println()

	windowMachine := &Windows{}
	windowMachineAdaptor := &WindowsAdaptor{
		windowMachine: windowMachine,
	}
	client.InsertLightningPortIntoComputer(windowMachineAdaptor)
}
