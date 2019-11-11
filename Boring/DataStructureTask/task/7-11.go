package task

import "math"

// 构造二叉链表的结构
type BiTree struct {
	Data int
	LChild *BiTree
	RChile *BiTree
}

var Nodesum = 0 //总结点数
var Leasum = 0	//总叶节点数

// 求二叉树的高度
func BTHigh(bt *BiTree) int{
	if bt == nil {//如果树为空，直接返回
		return 0
	}

	//否则分别遍历左右子树，高度为分别为左右子树加一
	LTree := BTHigh(bt.LChild)+1
	RTree := BTHigh(bt.RChile)+1

	//返回左右子树较大的一个
	return int(math.Max(float64(LTree), float64(RTree)))
}

// 求二叉树的宽度，利用层序遍历来求树的宽度
func BTWidth(bt *BiTree) int{
	r := []BiTree{*bt}	// 存储每层子树包含的节点
	max := len(r)	// 存储最大宽度
	var r1 []BiTree	//存储下一层的节点
	for  {
		for len(r) > 0 {//遍历该层所有的节点，将下一层的所有节点添加到 r1
			node := r[0]
			r = r[1:]
			if node.LChild != nil {
				r1 = append(r1, *node.LChild)
			}
			if node.RChile != nil {
				r1 = append(r1, *node.RChile)
			}
		}

		r = r1
		r1 = []BiTree{}
		// 如果当前层的宽度大于上一层的宽度，则更新 max 的值
		if len(r1) > max {
			max = len(r)
		}
		// 层序遍历结束，退出循环
		if len(r) == 0 {
			break
		}
	}

	return max
}

// 求二叉树的节点数
func BTPointSum(bt *BiTree) int{
	if bt == nil {
		return 0
	}
	Nodesum += BTPointSum(bt.LChild)+1
	Nodesum += BTPointSum(bt.RChile)+1
	return Nodesum
}

// 求二叉树的叶节点数
func BTTreePointSum(bt *BiTree) int{
	if bt.RChile == nil && bt.LChild == nil {
		return 1
	}
	Leasum += BTTreePointSum(bt.LChild)
	Leasum += BTTreePointSum(bt.RChile)
	return 0
}