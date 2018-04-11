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
