package main

import (
	"fmt"
	"sync"
)

type single struct {
}

var singleInstance *single
var mu = &sync.Mutex{}

func getInstance() *single {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}
