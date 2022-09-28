package array

import (
	"fmt"

	"github.com/testsabirweb/striver/utils"
)

func RotateMatrix() {
	arr := make([][]int, 7)
	arr[0] = []int{1, 2, 3, 4, 5, 6, 7}
	arr[1] = []int{8, 9, 10, 11, 12, 13, 14}
	arr[2] = []int{2, 3, 4, 5, 6, 7, 8}
	arr[3] = []int{1, 2, 3, 4, 5, 6, 7}
	arr[4] = []int{8, 9, 10, 11, 12, 13, 14}
	arr[5] = []int{2, 3, 4, 5, 6, 7, 8}
	arr[6] = []int{1, 2, 3, 4, 5, 6, 7}

	utils.Print2D[int](arr, len(arr), len(arr[0]))
	fmt.Println("######################################1")

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr[0]); j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	utils.Print2D[int](arr, len(arr), len(arr[0]))
	fmt.Println("######################################2")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0])/2; j++ {
			arr[i][j], arr[i][len(arr[0])-j-1] = arr[i][len(arr[0])-j-1], arr[i][j]
		}
	}
	utils.Print2D[int](arr, len(arr), len(arr[0]))
	fmt.Println("######################################3")
}
