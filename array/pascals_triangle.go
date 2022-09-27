package array

import "fmt"

func PascalsTriangle() {
	fmt.Println("Pascals Triangele")
	n := 5
	arr := make([]int, n+1)
	for a, _ := range arr {
		arr[a] = 1
	}
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			if i > 2 && j > 1 && j < i {
				arr[j] = arr[j] + arr[j-1]
			}
			fmt.Print(arr[j])
		}
		fmt.Println("")
	}
}

// 1
// 1 1
// 1 2 1
// 1 3 3 1
// 1 4 6 4 1
