package task

import "fmt"

// 用邻接表的存储结构实现图的基础操作
// 下面均使用无向图

// 边表结构
type EdgeNode struct {
	adjvex int	// 顶点的索引
	Next *EdgeNode	//指向下一邻接点
}

// 顶点表结构
type VertexNode struct {
	data string	// 顶点的数据
	firstEdge *EdgeNode	// 边表头指针
}

type Graph1 struct {
	AdjList []VertexNode
	numVertexes, numEdges int
}

// 在邻接表中插入一个新的顶点
// 首先需要将该点先添加到存储顶点元素的数组当中，其后的邻接点先置为空
// 然后依次更新邻接表的路径
func InsertVex1(G Graph1, v string){
	// 首先在邻接表中新增一项与 v 相关的空行
	vt := VertexNode{
		data:      v,
		firstEdge: nil,
	}
	G.AdjList = append(G.AdjList, vt)
	fmt.Printf("请输入加入 %s 后需要添加的线的条数：", v, "\n")
	var en int
	fmt.Scanln(&en)

	// 然后根据给定的路径添加到邻接表中
	for i := 0; i < en; i++ {
		fmt.Printf("请输入与 %s 相连的另一个顶点：")
		var ee string
		fmt.Scanln(&ee)
		InsertArc1(G, v, ee)
	}
	G.numVertexes+=1
}

// 在邻接表中插入一条新的边
// 确定边的索引之后，直接使用头插法插入
// 因为是无向图，所以要添加两条记录
func InsertArc1(G Graph1, v,w string){
	// 首先确定两个顶点的索引
	vl, wl := -1, -1
	for i := 0; i < G.numVertexes; i++ {
		if G.AdjList[i].data == v {
			vl = i
		} else if G.AdjList[i].data == w {
			wl = i
		}

		if vl != -1 && wl != -1 {
			break
		}
	}

	// 将两条边分别头插入对应的邻接表当中
	e1 := &EdgeNode{
		adjvex: vl,
		Next:   G.AdjList[vl].firstEdge,
	}
	G.AdjList[vl].firstEdge = e1

	e2 := &EdgeNode{
		adjvex: wl,
		Next:   G.AdjList[wl].firstEdge,
	}
	G.AdjList[wl].firstEdge = e2

	G.numEdges+=1

}

// 删除邻接表中的某个顶点
// 首先先删掉所有以 v 为路径端点的路径
// 然后将该端点从邻接表中删除
func DeleteVex1(G Graph1, v string){
	// 先删掉所有与 v 相连的线
	p := G.AdjList[0].firstEdge
	for p != nil {
		DeleteArc1(G, v, G.AdjList[p.adjvex].data)
		p = p.Next
	}

	// 然后将 v 从邻接表中移除
	if G.AdjList[len(G.AdjList)-1].data == v {
		G.AdjList = G.AdjList[:len(G.AdjList)-1]
	} else {
		for i := 0; i < G.numVertexes; i++ {
			if G.AdjList[i].data == v {
				G.AdjList = append(G.AdjList[0:i], G.AdjList[i+1:]...)
			}
		}
	}

	G.numVertexes-=1
}

// 删除邻接表中的一条边
// 确定边的两个端点的索引之后，直接将其从邻接表中删除
func DeleteArc1(G Graph1, v,w string){
	G.numEdges-=2
	// 首先确定 v，w 的索引
	vl, wl := -1, -1
	for i := 0; i < G.numVertexes; i++ {
		if G.AdjList[i].data == v {
			vl = i
		} else if G.AdjList[i].data == w {
			wl = i
		}

		if vl != -1 && wl != -1 {
			break
		}
	}

	DeleteNode(&G.AdjList[vl], wl)
	DeleteNode(&G.AdjList[wl], vl)
}

//
func DeleteNode(vertexNode *VertexNode, i int) {
	if vertexNode.firstEdge.adjvex == i && vertexNode.firstEdge.Next == nil {
		vertexNode.firstEdge = nil
		return
	} else if vertexNode.firstEdge.adjvex == i && vertexNode.firstEdge.Next != nil {
		vertexNode.firstEdge = vertexNode.firstEdge.Next
		return
	}

	p := vertexNode.firstEdge
	for p != nil {
		if p.Next.adjvex == i {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
	}
}