package utils

import (
	"fmt"
	"log"
	"os"
)

func GetCurrentPath() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}

type ListTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

// arr := []int{1, 2, 3, 4, 5}

// 	a := utils.Make2D[int](len(arr), len(arr))

// 	for i := 0; i < len(arr); i++ {
// 		a[i] = arr
// 	}

// 	utils.Print2D[int](a, len(a), len(a[0]))

func Print2D[T ListTypes](arr [][]T, n int, m int) {
	if n == 0 {
		n = len(arr)
	}
	if m == 0 {
		m = len(arr[0])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

// type stringer interface {
// 	String() string
// }

// func concat[T stringer](vals []T) string {
// 	result := ""
// 	for _, val := range vals {
// 		result += val.String()
// 	}
// 	return result
// }

func ReverseArray(a *[]int, start int, end int) {
	for start < end {
		(*a)[start], (*a)[end] = (*a)[end], (*a)[start]
		start++
		end--
	}
}
