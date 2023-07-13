package balance_binary_search_tree

func balanceBST(root *TreeNode) *TreeNode {
	values := []int{}
	dfs(root, &values)
	return build(values, 0, len(values)-1)
}

func dfs(node *TreeNode, values *[]int) {
	if node != nil {
		dfs(node.Left, values)
		*values = append(*values, node.Val)
		dfs(node.Right, values)
	}
}

func build(values []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2

	return &TreeNode{
		Val:   values[mid],
		Left:  build(values, left, mid-1),
		Right: build(values, mid+1, right),
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
