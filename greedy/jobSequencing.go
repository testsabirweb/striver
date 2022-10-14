package greedy

import (
	"fmt"
	"sort"
)

type job struct {
	name     string
	deadline int
	profit   int
}

type jobes []job

func (j jobes) Len() int {
	return len(j)
}

func (j jobes) Less(a, b int) bool {
	if j[a].deadline == j[b].deadline {
		return j[a].profit > j[b].profit
	}
	return j[a].deadline > j[b].deadline
}

func (j jobes) Swap(a, b int) {
	j[a], j[b] = j[b], j[a]
}

func Jobsequencing() {
	jobs := jobes{
		{"a", 2, 100},
		{"b", 1, 19},
		{"c", 2, 27},
		{"d", 1, 25},
		{"e", 3, 15},
	}
	fmt.Println(jobs)
	sort.Sort(jobes(jobs))

	fmt.Println(jobs)

	profit := 0
	deadline := jobs[0].deadline

	for i := 0; i < len(jobs); i++ {
		if deadline <= jobs[i].deadline && deadline > 0 {
			deadline--
			profit = profit + jobs[i].profit
			fmt.Println(jobs[i].name)
		}
	}

	fmt.Println(profit, "#############")

}
