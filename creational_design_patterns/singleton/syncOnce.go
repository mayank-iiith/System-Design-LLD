package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var singleInstanceUsingSyncOnce *single

func getInstanceUsingSyncOnce() *single {
	if singleInstanceUsingSyncOnce == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now.")
			singleInstanceUsingSyncOnce = &single{}
		})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstanceUsingSyncOnce
}
