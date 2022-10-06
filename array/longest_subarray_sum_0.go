package array

import "fmt"

// type location [][]int

func LongestSubarraySum0() {
	arr := []int{1, 2, 3}
	k := 3
	count := 0
	sum := arr[0]
	leftSum := map[int]int{}
	leftSum[arr[0]] = 1
	if sum == k {
		count++
	}
	fmt.Println(count, leftSum, sum)
	for i := 1; i < len(arr); i++ {
		sum = sum + arr[i]
		if sum == k {
			count++
		}
		_, exist := leftSum[sum-k]
		if exist {
			count++
		}
		leftSum[sum] = leftSum[sum] + 1
		fmt.Println(count, leftSum, sum)

	}
	fmt.Println(count)
}
