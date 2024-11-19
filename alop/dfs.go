package main

//
//import "fmt"
//
//// Graph 使用邻接表表示
//type Graph struct {
//	nodes   map[int][]int // key为节点，value为邻接节点列表
//	visited map[int]bool  // 记录节点是否被访问过
//}
//
//// AddEdge 添加边
//func (g *Graph) AddEdge(u, v int) {
//	g.nodes[u] = append(g.nodes[u], v) // 添加从u到v的边
//	// 如果是无向图，还需要添加从v到u的边
//	// g.nodes[v] = append(g.nodes[v], u)
//}
//
//// NewGraph 创建一个新的图实例
//func NewGraph() *Graph {
//	return &Graph{
//		nodes:   make(map[int][]int),
//		visited: make(map[int]bool),
//	}
//}
//
//// DFS 深度优先搜索
//func (g *Graph) DFS(start int) {
//	g.dfs(start)
//}
//
//// dfs 递归进行深度优先搜索
//func (g *Graph) dfs(node int) {
//	if g.visited[node] { // 如果已经访问过，则返回
//		return
//	}
//
//	g.visited[node] = true // 标记为已访问
//	fmt.Println(node)      // 打印节点
//
//	for _, n := range g.nodes[node] { // 遍历所有邻接节点
//		g.dfs(n)
//	}
//}
//
//func main3() {
//	graph := NewGraph()
//	// 假设我们有一个简单的图
//	graph.AddEdge(0, 1)
//	graph.AddEdge(0, 2)
//	graph.AddEdge(1, 3)
//	graph.AddEdge(1, 4)
//	graph.AddEdge(2, 5)
//	graph.AddEdge(2, 6)
//
//	fmt.Println("Starting DFS from node 0:")
//	graph.DFS(0)
//}
