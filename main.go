package main

import (
	"fmt"

	stackandqueue "github.com/testsabirweb/striver/stack_and_queue"
)

func main() {
	fmt.Println("Runnning main file")
	// array.SetMatrixZero()
	// array.PascalsTriangle()
	// fmt.Println(array.SellBuy([]int{7, 1, 5, 3, 6, 4}))
	// array.MergeOverlapping()

	// array.MajorityNBy2()
	// arr := []int{3, 5, 7, 2, 1, 9, 4, 8, 4, 6}
	// fmt.Println(algo.MergeSort(arr))
	// greedy.Jobsequencing()
	// recursion.Subset1()
	// recursion.CombinationSum1()

	// binarysearch.KthElementIn2Sorted()
	// heap.Heap()
	// stackandqueue.QueueUsingStack()
	// stackandqueue.NextGreaterElement()
	// stackandqueue.LRUCache()
	var arr = []int{5, 7, 1, 7, 6, 0}
	stackandqueue.LargestRect(arr)

}
