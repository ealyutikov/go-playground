package balance_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_balanceBST(t *testing.T) {
	assert.Equal(
		t,
		&TreeNode{},
		balanceBST(&TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		}),
	)
}
