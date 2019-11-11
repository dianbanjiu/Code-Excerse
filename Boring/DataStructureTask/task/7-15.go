package task

import "fmt"

// 使用邻接矩阵存储结构实现图的基本操作
// 以下使用的是无向图

type Graph struct {
	Vex                    []string
	Edge                   [][]int
	NumVertexes, NumEdgees int
}

// 添加新的点需要添加对应的行和列
func InsertVex(G Graph, v string) {
	G.NumVertexes += 1       // 将顶点数加一
	G.Vex = append(G.Vex, v) // 将新的顶点添加到点集
	// 为边集添加新的一行初始为 0 的行，并将原式数组长度加一
	ve := make([]int, G.NumVertexes)
	for i := 0; i < len(G.Edge); i++ {
		G.Edge[i] = append(G.Edge[i], 0)
	}
	// 输入与 v 相关的边的条数，并根据给定的下标将其插入到新的边集当中
	G.Edge = append(G.Edge, ve)
	fmt.Println("请输入需要添加的边的条数： ")
	var e int
	fmt.Scanln(&e)
	G.NumEdgees += e
	for j := 0; j < e; j++ {
		var vi string
		fmt.Printf("请输入与 %s 相关的点的名称", v, "\n")
		fmt.Scanln(&vi)
		InsertArc(G, v, vi)
	}
	var input string
	fmt.Println("请输入与")
}

// 在图中插入一条新的边，
// 将对应边在边集中的值设为 1 即可
func InsertArc(G Graph, v, w string) {
	G.NumEdgees += 1
	// 确定弧头 弧尾在点集中的索引
	vl, wl := -1, -1
	for i := 0; i < G.NumVertexes; i++ {
		if G.Vex[i] == v {
			vl = i
		} else if G.Vex[i] == w {
			wl = i
		}

		if vl != -1 && wl != -1 {
			break
		}
	}

	// 从边集中找到对应的边，然后将其值置为 1
	for i := 0; i < G.NumVertexes; i++ {
		for j := 0; j < G.NumVertexes; j++ {
			if i == vl && j == wl {
				G.Edge[i][j] = 1
			} else if i == wl && j == vl {
				G.Edge[i][j] = 1
			}
		}
	}
}
func DeleteVex(G Graph, v string) {
	G.NumVertexes -= 1
	var local int
	for i := 0; i < G.NumVertexes; i++ { //查找 v 在点集中的位置
		if G.Vex[i] == v {
			local = i
		}
	}

	// 将边的二维数组中 v 所在的行去掉
	if local != G.NumVertexes-1 {
		G.Edge = append(G.Edge[0:local], G.Edge[local+1:]...)
	} else {
		G.Edge = G.Edge[0:local]
	}

	// 遍历整个数组，将 v 索引所在的列从边数组中去掉
	for i := 0; i < len(G.Edge); i++ {
		var temp []int
		for j := 0; j < G.NumVertexes; j++ {
			if i != local && j != local {
				temp = append(temp, G.Edge[i][j])
			}
			G.Edge[i] = temp
		}
	}
}

// 删除有向图的一条弧，v 为弧尾 w 为 弧头
func DeleteArc(G Graph, v, w string) {
	G.NumEdgees -= 1
	// 确定弧头 弧尾在点集中的索引
	vl, wl := -1, -1
	for i := 0; i < G.NumVertexes; i++ {
		if G.Vex[i] == v {
			vl = i
		} else if G.Vex[i] == w {
			wl = i
		}

		if vl != -1 && wl != -1 {
			break
		}
	}

	// 从边集中找到对应的边，然后将其值置为 0
	for i := 0; i < G.NumVertexes; i++ {
		for j := 0; j < G.NumVertexes; j++ {
			if i == vl && j == wl {
				G.Edge[i][j] = 0
			} else if i == wl && j == vl {
				G.Edge[i][j] = 0
			}
		}
	}
}
