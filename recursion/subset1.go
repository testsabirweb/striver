package recursion

import "fmt"

func printSum(arr []int) {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	fmt.Println(sum)
}

func getSum(arr []int) {
	// printSum(arr)
	fmt.Println(arr)

	for i := 0; i < len(arr) && len(arr) > 2; i++ {
		temp := []int{}
		for j := 0; j < len(arr); j++ {
			if i != j {
				temp = append(temp, arr[j])
			}
		}

		getSum(temp)
		// if i >= len(arr)-1 {
		// 	copy(temp, arr[:i])
		// 	getSum(temp)
		// } else {
		// 	copy(temp, append(arr[0:i], arr[i+1:]...))
		// 	getSum(temp)
		// }
	}

}

func Subset1() {
	arr := []int{5, 2, 1}
	getSum(arr)
}
