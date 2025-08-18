package main

// Decorator is a structural pattern that allows adding new behaviors to objects dynamically by placing them inside special wrapper objects, called decorators.

// Using decorators you can wrap objects countless number of times since both target objects and decorators follow the same interface. The resulting object will get a stacking behavior of all wrappers.

import (
	"fmt"
	datasource "lld/structural_design_patterns/decorator/data_source"
	"log"
	"os"
)

func main() {
	salaryRecords := `Name, Salary
	John Smith,100000
	Steven Jobs,912000`

	fileName := "output.txt"

	// Stack: Compression -> Encryption -> File

	wrapped := datasource.NewCompressionDecorator(
		datasource.NewEncryptionDecorator(
			datasource.NewFileDataSource(fileName),
		),
	)

	if err := wrapped.WriteData(salaryRecords); err != nil {
		log.Fatal("error while writing:", err)
	}

	plain := datasource.NewFileDataSource(fileName)
	compressedAndEncoded, err := plain.ReadData()
	if err != nil {
		log.Fatal("error while reading:", err)
	}

	decompressedAndDecoded, err := wrapped.ReadData()
	if err != nil {
		log.Fatal("error while reading:", err)
	}

	fmt.Println("- Input ----------------")
	fmt.Println(salaryRecords)
	fmt.Println()

	fmt.Println("- Compressed & Encoded --------------")
	fmt.Println(compressedAndEncoded)
	fmt.Println()

	fmt.Println("- Decompressed & Decoded --------------")
	fmt.Println(decompressedAndDecoded)

	os.Remove(fileName)
}
