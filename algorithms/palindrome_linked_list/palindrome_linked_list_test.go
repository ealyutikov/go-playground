package palindrome_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome_1(t *testing.T) {
	tree := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	assert.True(t, isPalindrome(tree))
}

func TestIsPalindrome_2(t *testing.T) {
	tree := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	assert.True(t, isPalindrome(tree))
}

func TestIsPalindrome_3(t *testing.T) {
	tree := &ListNode{1, &ListNode{2, nil}}
	assert.False(t, isPalindrome(tree))
}
