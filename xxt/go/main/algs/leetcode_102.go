package main

import "fmt"

/*
*

		给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
		输入：root = [3,9,20,null,null,15,7]
		输出：[[3],[9,20],[15,7]]
	 * Definition for a binary tree node.
	 * type TreeNode struct {
	 *     Val int
	 *     Left *TreeNode
	 *     Right *TreeNode
	 * }
		示例 2：

		输入：root = [1]
		输出：[[1]]
		示例 3：

		输入：root = []
		输出：[]

		提示：

		树中节点数目在范围 [0, 2000] 内
		-1000 <= Node.val <= 1000
*/
func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	traversalResult := levelOrder(root)
	fmt.Println(traversalResult)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var nextLevelNodes []*TreeNode
		var currentLevelValues []int
		for _, node := range nodes {
			currentLevelValues = append(currentLevelValues, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		result = append(result, currentLevelValues)
		nodes = nextLevelNodes
	}
	return result
}

/**
1. 定义TreeNode结构体，包含一个整数Value和两个指向TreeNode的指针Left和Right。这个结构体表示二叉树的节点。
2. 定义levelOrder函数，接收一个指向TreeNode的指针root作为输入参数。这个函数的目的是返回一个嵌套切片，包含二叉树的层次遍历结果。
3. 检查root是否为nil。如果是，则返回nil，因为这意味着没有输入树。
4. 初始化一个名为result的空切片。这个切片将用于存储最终的层次遍历结果。
5. 创建一个名为nodes的切片，初始值为包含根节点的切片。这个切片将用于迭代地遍历每一层的节点。
6. 使用for循环，只要nodes切片非空，就执行以下操作：
	a. 初始化一个名为nextLevelNodes的空切片。这个切片将用于存储下一层的节点。
	b. 初始化一个名为currentLevelValues的空切片。这个切片将用于存储当前层的节点值。
	c. 使用for循环，遍历nodes切片中的每个节点：
		i. 将当前节点的值添加到currentLevelValues切片。
		ii. 如果当前节点有左子节点，将左子节点添加到nextLevelNodes切片。
		iii. 如果当前节点有右子节点，将右子节点添加到nextLevelNodes切片。
	d. 将currentLevelValues切片添加到result切片。
	e. 更新nodes切片为nextLevelNodes切片，以便在下次迭代中处理下一层的节点。
7. 循环结束后，返回result切片，它包含了整个二叉树的层次遍历结果。
8. 在main函数中，创建一个示例二叉树。
9. 调用levelOrder函数，传入示例二叉树的根节点，并将返回的嵌套切片赋值给traversalResult。
10. 打印traversalResult，它将按层次顺序显示二叉树的节点值。


1. 确定要实现的功能：层次遍历二叉树，并返回一个嵌套的整数切片。
2. 使用一个切片nodes来存储每一层的节点。在循环的每次迭代中，我们将处理这个切片中的所有节点。
3. 在循环的每次迭代中，我们创建一个名为currentLevelValues的切片来存储当前层的节点值，以及一个名为nextLevelNodes的切片来存储下一层的节点。
4. 对于nodes切片中的每个节点，我们将其值添加到currentLevelValues切片，并将其子节点（如果存在）添加到nextLevelNodes切片。
5. 将currentLevelValues切片添加到result切片，表示已经完成了对当前层的处理。
6. 更新nodes切片为nextLevelNodes切片，以便在下次循环迭代中处理下一层的节点。
7. 当nodes切片为空时，表示已经完成了对所有层的处理，此时退出循环。
8/ 返回result切片，它包含了整个二叉树的层次遍历结果。

这种实现方法的关键在于，我们使用了一个切片nodes来追踪每一层的节点。
在处理完当前层的节点后，我们使用nextLevelNodes切片收集下一层的节点，并更新nodes切片，
以便在下次循环迭代中处理下一层的节点。这样，我们可以避免使用队列，但仍然能够实现层次遍历二叉树的功能。

*/
