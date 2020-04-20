/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    t := &TreeNode{Val: preorder[0]}
    for _, val := range preorder[1:] {
        InsertNode(t, val)
    }
    return t
}


func InsertNode(t *TreeNode, value int) {
    if value < t.Val {
        if t.Left == nil {
            t.Left = &TreeNode{Val: value}
            return
        }
        InsertNode(t.Left, value)
    }
    if value > t.Val {
        if t.Right == nil {
            t.Right = &TreeNode{Val: value}
            return
        }
        InsertNode(t.Right, value)
    }
}