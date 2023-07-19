package lru_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)

	assert.Equal(
		t,
		cache.Get(666),
		-1,
	)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	assert.Equal(
		t,
		cache.Get(3),
		3,
	)
}
