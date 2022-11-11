package binarysearch

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func average(a, b int) int {
	return (a + b) / 2
}

func KthElementIn2Sorted() {
	array1 := []int{1, 4, 8, 10}
	array2 := []int{2, 3, 6, 7, 9}
	k := 5
	result := 0

	big := make([]int, len(array1))
	small := make([]int, len(array2))

	if len(array1) > len(array2) {
		copy(big, array1)
		copy(small, array2)
	} else {
		copy(big, array2)
		copy(small, array1)
	}
	smallmid := k / 2
	bigmid := k / 2

	bl := smallmid - 1
	br := smallmid
	sl := bigmid - 1
	sr := bigmid

	if big[bl] <= small[sr] && big[br] >= small[sl] {
		result = min(small[sr], big[br])
	} else if big[bl] > small[sr] {
		bigmid = bigmid - average(smallmid, len(small))
		smallmid = smallmid + average(smallmid, len(small))

	} else {
		bigmid = bigmid + average(0, smallmid)
		smallmid = smallmid - average(0, smallmid)
	}

	fmt.Println(result)

}
