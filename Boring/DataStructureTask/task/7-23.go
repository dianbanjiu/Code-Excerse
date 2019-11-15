package task

// 基于广度优先算法判断两顶点之间是否存在路径
// 以邻接表为存储结构
// 解法：从路径头开始对图进行广度搜索，直至找到队尾顶点，否则返回 false
func BFSTraverse(G Graph1, v,w string) bool{
	var r bool	// 存储返回结果
	visited := make([]bool, G.numVertexes)	// 用来判断顶点是否已经被访问的数组，初始化为 false
	vi = local(G, v)	// 路径起点的编号
	wi = local(G, w)	// 路径终点的编号
	var queue []int	// 存储每层节点的队列
	queue = append(queue, vi)	// 首先将路径起点加入队列

	// 若队列非空，则进行下一步
	for len(queue) != 0 {
		u := queue[0]	// 取出队顶元素
		queue = queue[1:]	// 缩短队列

		// 对每个节点相邻的顶点进行遍历，若该顶点未访问过，则将其添加到队尾
		for ue := G.AdjList[u].firstEdge; ue != nil; ue = ue.Next {
			if !visited[ue.adjvex] {
				visited[ue.adjvex]=true
				if ue.adjvex == wi {	// 若此时访问到的顶点为路径尾顶点，结束循环，返回肯定的判断结果
					r = true
					break
				}
				queue = append(queue, ue.adjvex)
			}
		}
	}
	return r
}