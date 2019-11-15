package task

// 基于图的深度优先算法判断两个顶点之间是否存在路径
// 以邻接表为存储结构
var (
	visited []bool	// 节点是否被访问的数组
	vi, wi int	// 存储路径首尾的编号
	r bool	// 存储返回结果
)

func DFSTraverse(G Graph1, v,w string) bool{
	visited = make([]bool, G.numVertexes)	// 初始化所有节点的访问情况为 false
	vi = local(G, v)	// 确定路径头的编号
	wi = local(G, w)	// 确定路径尾的编号
	DFS(G, vi)	// 以路径头为起点进行深度优先搜索
	return r
}

func DFS(G Graph1, v int) {
	visited[v] = true	// 将访问过得节点标记
	for w := G.AdjList[v].firstEdge;w!= nil ;w= w.Next  {// 对每一条连通路径进行访问标记
		if visited[w.adjvex] {
			DFS(G, w.adjvex)
		}
		if w.adjvex == wi {	// 直至找到对应的顶点，此时中断遍历将返回结果标记为 true
			r=true
			break
		}
	}
}

// 定位 s 在图中的编号
func local(G Graph1,s string) int {
	var i int
	for i = 0; i< G.numVertexes; i++  {
		if G.AdjList[i].data == s {
			return i
		}
	}
	return i
}