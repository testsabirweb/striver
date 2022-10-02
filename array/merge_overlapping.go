package array

import (
	"fmt"
	"sort"
)

type byEnd [][]int

func (s byEnd) Len() int {
	return len(s)
}
func (s byEnd) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byEnd) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func MergeOverlapping() {
	arr := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// fmt.Println(arr)
	sort.Sort(byEnd(arr))
	// fmt.Println(arr)

	result := [][]int{}

	temp := arr[0][:]
	for i := 1; i < len(arr); i++ {
		if arr[i-1][1] > arr[i][0] {
			temp[1] = arr[i][1]
		} else {
			result = append(result, temp)
			temp = arr[i][:]
		}
	}
	fmt.Println(result)

}
