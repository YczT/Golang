package main

import "fmt"

/*
节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。


示例1:
 输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]],
start = 0, target = 2

 输出：true

示例2:
 输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4],
[0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]],
start = 0, target = 4

 输出 true
*/

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	g := make([][]int, n) //建立一个二维数组，用来存放以所有开始所能到达的位置
	for i := 0; i < n; i++ {
		g[i] = []int{}
	}
	for i := 0; i < len(graph); i++ {
		st, ed := graph[i][0], graph[i][1]
		g[st] = append(g[st], ed)
	}
	fmt.Printf("%v\n", g)
	return DFS(g, start, target)
}

func DFS(g [][]int, st int, ed int) bool {
	for _, elem := range g[st] {
		if elem == ed {
			return true
		}
		//fmt.Printf("g:%v\n", g)
		temp := DFS(g, elem, ed)
		if temp {
			return true
		}
	}
	return false
}

func main() {
	var g = [][]int{{0, 1}, {0, 2}, {0, 4}, {0, 4}, {0, 1}, {1, 3}, {1, 4}, {1, 3}, {2, 3}, {3, 4}}
	//fmt.Printf("%v\n", g[0])
	//fmt.Printf("%v\n", len(g))
	ans := findWhetherExistsPath(5, g, 0, 4)
	fmt.Printf("结果:%v\n", ans)

}

/*
//输出：
[[1 2 4 4 1] [3 4 3] [3] [4] []]
结果:true

*/
