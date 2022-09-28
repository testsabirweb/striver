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

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func ReverseArray(a *[]int, start int, end int) {
	for start < end {
		(*a)[start], (*a)[end] = (*a)[end], (*a)[start]
		start++
		end--
	}
}

