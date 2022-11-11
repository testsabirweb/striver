package recursion

import "fmt"

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func solve(arr []int, target int, index int, result []int, fi [][]int) {

	if target == 0 {
		fi = append(fi, result)
	}
	if index < len(arr) {
		resultCopy := make([]int, len(result))
		copy(resultCopy, result)
		resultCopy = append(resultCopy, arr[index])
		solve(arr, target, index+1, resultCopy, fi)
		resultCopy2 := make([]int, len(result))
		copy(resultCopy2, result)
		solve(arr, target, index+1, resultCopy2, fi)
	}
}

func CombinationSum1() {
	arr := []int{2, 3, 6, 7}
	target := 7

	result := []int{}
	fi := [][]int{}

	solve(arr, target, 0, result, fi)

	fmt.Println(fi)

}
