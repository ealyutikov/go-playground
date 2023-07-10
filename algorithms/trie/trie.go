package trie

type Trie struct {
	root *Node
	size int
}

type Node struct {
	children map[rune]*Node
	last     bool
}

func InitTrie() *Trie {
	return &Trie{root: &Node{make(map[rune]*Node), false}}
}

func (t *Trie) Put(word string) bool {
	exists := true
	currentNode := t.root

	for _, letter := range word {
		n, ok := currentNode.children[letter]
		if !ok {
			exists = false
			n = &Node{make(map[rune]*Node), false}
			currentNode.children[letter] = n
		}
		currentNode = n
	}

	currentNode.last = true

	if !exists {
		t.size++
	}

	return exists
}

func (t *Trie) Contains(word string) bool {
	node, _ := t.getNode(word)
	return node != nil && node.last
}

func (t *Trie) Search(prefix string) []string {
	node, r := t.getNode(prefix)
	return search(node, r, []rune(prefix[:len(prefix)-1]))
}

func search(currentNode *Node, currentRune rune, prefix []rune) []string {
	result := []string{}

	newPrefix := append(prefix, currentRune)
	if currentNode.last {
		result = append(result, string(newPrefix))
	}

	for letter, node := range currentNode.children {
		newWords := search(node, letter, newPrefix)
		result = append(result, newWords...)
	}

	return result
}

func (t *Trie) getNode(prefix string) (*Node, rune) {
	currentNode := t.root
	var r rune
	for _, letter := range prefix {
		n, ok := currentNode.children[letter]
		if !ok {
			return nil, 0
		}
		currentNode = n
		r = letter
	}

	return currentNode, r
}
