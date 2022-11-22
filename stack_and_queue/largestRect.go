package stackandqueue

import "fmt"

func area(left []int, right []int, index int, height int) int {
	return (right[index] - left[index] - 1) * height
}

func LargestRect(arr []int) {

	right := NextSmaller(arr, true)
	left := NextSmaller(arr, false)
	result := 0
	li := 0
	ri := 0
	for i := 0; i < len(right); i++ {
		if result < area(left, right, i, arr[i]) {
			result = area(left, right, i, arr[i])
			li = left[i]
			ri = right[i]
		}
	}
	fmt.Println(result, li, ri)

}
