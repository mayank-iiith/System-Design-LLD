package cache

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	size         int
}

func NewCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		capacity:     2,
		size:         0,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) Add(key, value string) {
	if c.size == c.capacity {
		c.evict()
	}
	c.size++
	c.storage[key] = value
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.size--
}
