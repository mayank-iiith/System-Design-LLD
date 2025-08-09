package cache

import "fmt"

type LFUEvictionAlgo struct {
}

func (e *LFUEvictionAlgo) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy.")
}
