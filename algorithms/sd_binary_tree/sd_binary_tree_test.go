package sd_binary_tree

import (
	"fmt"
	"testing"
)

func TestBTree(t *testing.T) {
	codec := Constructor()
	string1 := codec.serialize(&TreeNode{
		1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	})

	string2 := codec.serialize(nil)

	fmt.Printf("%s\n", string1)
	fmt.Printf("%s", string2)
}
