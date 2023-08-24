package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	stack := Stack{}

	node := head
	for node != nil {
		stack.Push(node.Val)
		node = node.Next
	}

	node = head
	for node != nil {
		e, _ := stack.Pop()
		if e != node.Val {
			return false
		}

		node = node.Next
	}

	return stack.IsEmpty()
}

type Stack []int

func (s *Stack) Push(elem int) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	topIndex := len(*s) - 1
	topElem := (*s)[topIndex]
	*s = (*s)[:topIndex]
	return topElem, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
