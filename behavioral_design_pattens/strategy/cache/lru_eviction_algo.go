package cache

import "fmt"

type LRUEvictionAlgo struct {
}

func (e *LRUEvictionAlgo) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy.")
}
