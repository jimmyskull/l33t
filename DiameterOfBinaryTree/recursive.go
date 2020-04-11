// subtreeDiameter iterate over nodes of a tree and compute the
// longest path between any two leaf nodes.
func subtreeDiameter(root *TreeNode, max *int) (height int) {
    // If the root of the current node, there are no paths in this
    // subtree and the diameter within this subtree is zero.
    if root == nil {
        height = 0
        return
    }
    var left int
    var right int
    // Compute the diameter of left and right subtrees and increase
    // its height by one due to the new edge.
    // The max variable is passed by pointer and will be updated 
    // within the subtrees.
    if root.Left != nil {
        left = subtreeDiameter(root.Left, max) + 1
        
    }
    if root.Right != nil {
        right = subtreeDiameter(root.Right, max) + 1
    }
    // Compute the diameter of the current subtree and update the
    // global maximum.
    currMax := left + right
    if currMax > *max {
        *max = currMax
    }
    // Return the height that is the maximum between the left and
    // right subtrees.
    if left > right {
        height = left
    } else {
        height = right
    }
    return
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var diameter int
    subtreeDiameter(root, &diameter)
    return diameter
}