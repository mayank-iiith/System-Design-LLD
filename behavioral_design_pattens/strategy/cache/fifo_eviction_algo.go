package cache

import "fmt"

type FIFOEvictionAlgo struct {
}

func (e *FIFOEvictionAlgo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy.")
}
