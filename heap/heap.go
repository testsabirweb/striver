package heap

import "fmt"

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func heapify(arr []int, index int) {
	largestIndex := index
	left, right := GetChildNodeInHeap(index)

	if left < len(arr) && arr[left] > arr[largestIndex] {
		largestIndex = left
	}
	if right < len(arr) && arr[right] > arr[largestIndex] {
		largestIndex = right
	}

	if largestIndex != index {
		swap(arr, index, largestIndex)

		heapify(arr, largestIndex)
	}

}

func GetNodeParentInHeap(index int) int {
	return (index - 1) / 2
}
func GetChildNodeInHeap(index int) (int, int) {
	return 2*index + 1, 2*index + 2
}

func GenerateHeap(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i)
	}
	return arr
}

func removeMax(arr []int, lastIndex int) {
	swap(arr, 0, lastIndex)
	heapify(arr, 0)
}

func heapSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		removeMax(arr, i)
	}
}

func Heap() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println(GenerateHeap(arr))
	heapSort(arr)
	fmt.Println(arr)

}
