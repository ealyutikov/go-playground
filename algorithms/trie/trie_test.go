package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := InitTrie()
	trie.Put("apple")
	trie.Put("applo")
	trie.Put("car")

	assert.Equal(t, 3, trie.size)
	assert.False(t, trie.Contains("test"))
	assert.True(t, trie.Contains("apple"))
	assert.ElementsMatch(t, []string{"apple", "applo"}, trie.Search("app"))
}
