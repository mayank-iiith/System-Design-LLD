package cache

type EvictionAlgo interface {
	evict(c *Cache)
}
