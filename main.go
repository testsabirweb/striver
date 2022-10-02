package main

import (
	"fmt"

	"github.com/testsabirweb/striver/algo"
)

func main() {
	fmt.Println("Runnning main file")
	// array.SetMatrixZero()
	// array.PascalsTriangle()
	// fmt.Println(array.SellBuy([]int{7, 1, 5, 3, 6, 4}))
	// array.MergeOverlapping()

	// array.MajorityNBy2()
	arr := []int{3, 5, 7, 2, 1, 9, 4, 8, 4, 6}
	fmt.Println(algo.MergeSort(arr))
}
