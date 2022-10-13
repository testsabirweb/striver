package greedy

import (
	"fmt"
	"sort"
)

type matrix [][]int

func (m matrix) Len() int {
	return len(m)
}

func (m matrix) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m matrix) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}

func NMeeting1Room() {
	n := 6
	start := []int{1, 3, 0, 5, 8, 5}
	end := []int{2, 4, 5, 7, 9, 9}
	comp := [][]int{}
	count := 1
	for i := 0; i < n; i++ {
		comp = append(comp, []int{start[i], end[i]})
	}
	// fmt.Println(comp)
	sort.Sort(matrix(comp))
	// fmt.Println(comp)

	for i := 1; i < n; i++ {
		if comp[i][0] >= comp[i-1][1] {
			count++
		}
	}

	fmt.Println(count)
}
