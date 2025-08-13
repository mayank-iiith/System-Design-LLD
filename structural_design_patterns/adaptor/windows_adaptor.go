package main

import "fmt"

type WindowsAdaptor struct {
	windowMachine *Windows
}

func (w *WindowsAdaptor) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.InsertIntoUSBPort()
}
