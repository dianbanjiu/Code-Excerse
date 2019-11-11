package task

// 基于广度优先算法判断两顶点之间是否存在路径
// 以邻接表为存储结构

var (
	visited []bool
	vi, wi int
	r bool
)

func DFSTraverse(G Graph1)