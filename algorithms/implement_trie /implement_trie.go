package implement_trie

type Trie struct {
	root *Node
}

type Node struct {
	children map[rune]*Node
	last     bool
}

func Constructor() Trie {
	return Trie{root: &Node{make(map[rune]*Node), false}}
}

func (this *Trie) Insert(word string) {
	currentNode := this.root
	for _, letter := range word {
		n, ok := currentNode.children[letter]
		if !ok {
			n = &Node{make(map[rune]*Node), false}
			currentNode.children[letter] = n
		}
		currentNode = n
	}

	currentNode.last = true
}

func (this *Trie) Search(word string) bool {
	node := this.getNode(word)
	return node != nil && node.last
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.getNode(prefix)
	return node != nil && !node.last
}

func (t *Trie) getNode(prefix string) *Node {
	currentNode := t.root
	for _, letter := range prefix {
		n, ok := currentNode.children[letter]
		if !ok {
			return nil
		}
		currentNode = n
	}

	return currentNode
}
