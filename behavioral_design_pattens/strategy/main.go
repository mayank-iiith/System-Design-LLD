package main

// Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.
// The original object, called context, holds a reference to a strategy object. The context delegates executing the behavior to the linked strategy object. In order to change the way the context performs its work, other objects may replace the currently linked strategy object with another one.
// Ref: https://refactoring.guru/design-patterns/strategy/go/example#example-0

import (
	"lld/behavioral_design_pattens/strategy/cache"
)

func main() {
	fifoEvictionAlgo := &cache.FIFOEvictionAlgo{}
	cache1 := cache.NewCache(fifoEvictionAlgo)

	cache1.Add("a", "1")
	cache1.Add("b", "2")

	cache1.Add("c", "3")

	lruEvictionAlgo := &cache.LRUEvictionAlgo{}
	cache1.SetEvictionAlgo(lruEvictionAlgo)
	cache1.Add("d", "4")

	lfuEvictionAlgo := &cache.LFUEvictionAlgo{}
	cache1.SetEvictionAlgo(lfuEvictionAlgo)
	cache1.Add("e", "5")
}
