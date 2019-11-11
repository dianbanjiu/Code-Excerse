package task

// 题目：给定一个二叉树的前中序，构建该二叉树的二叉链表

// 定义二叉链表的节点结构
type BinTreeNode struct {
	Data string
	LChild, RChild *BinTreeNode
}

// 根据二叉给定的前中序数组确定创建二叉链表
// 算法结构：根据先序数组找到根节点，
// 根据根节点将中序数组分为左右子树。
// 对左右子树分别重复执行上面的操作

func CreateTree(T *BinTreeNode, pre, mid []string){
	if len(mid) == 0{
		T = nil
		return
	}

	T.Data = pre[0]	//创建根节点的数据
	index := 0	// 存储根节点的索引
	for i, v := range mid{
		if pre[0] == v{
			index = i
		}
	}

	CreateTree(T.LChild, pre[1:index], mid[:index])	// 递归创建左子树
	CreateTree(T.RChild, pre[index+1:], mid[index+1:])	// 递归创建右子树
}