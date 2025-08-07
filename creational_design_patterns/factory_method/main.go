package main

import (
	"fmt"
	"lld/creational_design_patterns/factory_method/storage"
	"log"
)

func main() {
	// store, err := storage.NewStorage(storage.DiskStorageType)
	// store, err := storage.NewStorage(storage.StorageType(5))

	store, err := storage.NewStorage(storage.MemoryStorageType)
	if err != nil {
		log.Fatalf("error while getting storage: %+v", err)
	}

	err = store.SaveData("hello, world!")
	if err != nil {
		log.Fatalf("error while saving data: %+v", err)
	}

	data, err := store.GetData()
	if err != nil {
		log.Fatalf("error while getting data: %+v", err)
	}

	fmt.Println("Data:", data)
}
