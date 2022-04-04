package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const capacity = 3
const value = 1 << 31

func TestCache(t *testing.T) {
	var cache = Constructor(capacity)

	// fill cache
	for key := 1; key <= capacity; key++ {
		cache.Put(key, value)
	}

	for key := 1; key <= capacity; key++ {
		v := cache.Get(key)
		assert.Equal(t, v, value)
	}

	// Trigger LRU
	cache.Put(capacity+1, value)
	v := cache.Get(1)
	assert.Equal(t, v, -1)

	// Trigger LRU, make key "2" not be eliminated
	_ = cache.Get(2)
	cache.Put(capacity+2, value)
	v = cache.Get(3)
	assert.Equal(t, v, -1)
}
