package main

import "fmt"

func main() {
	for range 50 {
		go getInstance()
		// go getInstanceUsingSyncOnce()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
