package sd_binary_tree

import "encoding/json"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const EmptyNode = "null"

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return EmptyNode
	}

	bytes, _ := json.Marshal(root)
	return string(bytes)
}

func (this *Codec) deserialize(data string) *TreeNode {
	if EmptyNode == data {
		return nil
	}

	node := TreeNode{}
	json.Unmarshal([]byte(data), &node)
	return &node
}
