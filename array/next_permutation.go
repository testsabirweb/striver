package array

import (
	"fmt"
	"test/workspace/src/github.com/testsabirweb/striver/utils"
)

func NextPermutation() {
	arr := []int{1, 2, 6, 8, 5, 4, 1}
	i := len(arr) - 1
	j := i

	for i > 0 {
		if arr[i-1] < arr[i] {
			i--
			break
		}
		i--
	}
	for j > i {
		if arr[j] > arr[i] {
			break
		}
		j--
	}

	if i == 0 {
		fmt.Println("It returns to case 54321")
	} else {
		arr[i], arr[j] = arr[j], arr[i]
		utils.ReverseArray(&arr, i+1, len(arr)-1)
		fmt.Println(arr)
	}
}
