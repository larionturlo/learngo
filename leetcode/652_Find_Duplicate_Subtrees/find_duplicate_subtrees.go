package main

import "fmt"
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mapSubtree(root *TreeNode, valueToNodeMap *map[string]bool, res *[]*TreeNode) string {
	if root == nil {
		return ""
	}
	key := fmt.Sprintf("%dl%sr%s", root.Val, mapSubtree(root.Left, valueToNodeMap, res), mapSubtree(root.Right, valueToNodeMap, res) )
	added, exists := (*valueToNodeMap)[key]
	if exists && !added {
		(*valueToNodeMap)[key] = true
		*res = append(*res, root)
	}
	if !exists {
		(*valueToNodeMap)[key] = false
	}
	return key
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	mapNodes := make(map[string]bool)
	res := make([]*TreeNode, 0, 0)
	mapSubtree(root, &mapNodes, &res)
	return res
}
