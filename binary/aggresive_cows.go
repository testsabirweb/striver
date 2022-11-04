package binary

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func AggresiveCows() {
	arr := []int{1, 2, 4, 8, 9}
	cows := 3
	result := 0
	sort.Ints(arr)
	low := arr[0]
	high := arr[len(arr)-1] - low

	for low <= high {

		mid := (low + high) / 2
		fmt.Println(low, mid, high)
		count := 1
		selectedIndex := 0
		for i := 1; i < len(arr); i++ {
			if arr[i]-arr[selectedIndex] >= mid {
				selectedIndex = i
				count++
			}
		}
		if count >= cows {
			result = max(result, mid)
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println(result)
}
