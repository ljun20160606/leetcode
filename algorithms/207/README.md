# Course Schedule

## 描述

```txt

There are a total of n courses you have to take, labeled from 0 to n - 1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]


Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?


For example:
2, [[1,0]]
There are a total of 2 courses to take. To take course 1 you should have finished course 0. So it is possible.

2, [[1,0],[0,1]]
There are a total of 2 courses to take. To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

Note:

The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.



click to show more hints.

Hints:

This problem is equivalent to finding if a cycle exists in a directed graph. If a cycle exists, no topological ordering exists and therefore it will be impossible to take all courses.
Topological Sort via DFS - A great video tutorial (21 minutes) on Coursera explaining the basic concepts of Topological Sort.
Topological sort could also be done via BFS.


```

## 题解

```go
package algorithms

func canFinish(numCourses int, prerequisites [][]int) bool {
	color := make([]int, numCourses)
	m := make(map[int][]int)
	for _, v := range prerequisites {
		i, i2 := v[0], v[1]
		if _, has := m[i]; !has {
			m[i] = []int{}
		}
		m[i] = append(m[i], i2)
	}
	isDAG := true
	var helper func(node int)
	helper = func(node int) {
		if !isDAG {
			return
		}
		color[node] = 1
		for _, v := range m[node] {
			switch color[v] {
			case 1:
				isDAG = false
				break
			case -1:
				continue
			default:
				helper(v)
			}
		}
		color[node] = -1
	}

	for k := range m {
		helper(k)
	}
	return isDAG
}

```

```go
package algorithms

func canFinish(numCourses int, prerequisites [][]int) bool {
	in := make([]int, numCourses)
	edge := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		src, dest := prerequisites[i][1], prerequisites[i][0]
		in[dest]++
		edge[src] = append(edge[src], dest)
	}
	vertexTraversed := numCourses
	queue := make([]int, 0, numCourses/2)
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		vertexTraversed--
		front := queue[0]
		queue = queue[1:]
		for i := 0; i < len(edge[front]); i++ {
			in[edge[front][i]]--
			if in[edge[front][i]] == 0 {
				queue = append(queue, edge[front][i])
			}
		}
		edge[front] = []int{}
	}
	if vertexTraversed != 0 {
		return false
	}
	return true
}

```