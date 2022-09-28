package array

import (
	"fmt"
)

func Kdane() int {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	i := 0
	j := 0
	start := 0
	end := 0
	k := 0
	currK := 0

	if len(arr) < 1 {
		return -1
	} else if len(arr) == 1 {
		return arr[0]
	}
	for {
		if i >= len(arr) || j >= len(arr) {
			break
		}
		currK = currK + arr[j]
		if currK > k {
			k = currK
			start = i
			end = j
		} else if currK < 0 {
			i = j
			currK = 0
		}
		j++

	}
	fmt.Println(start+1, end, k)

	return k
}
