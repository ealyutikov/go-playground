package implement_trie

type Trie struct {
	nodes []rune
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	return
}

func (this *Trie) Search(word string) bool {
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
