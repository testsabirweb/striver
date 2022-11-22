package stackandqueue

import "fmt"

func NextSmaller(arr []int, direction bool) []int {
	// var arr = []int{5, 7, 1, 7, 6, 0}
	arrCopy := make([]int, len(arr))
	if !direction {
		for i := 0; i < len(arr); i++ {
			arrCopy[i] = len(arr) - 1
		}
	}

	s := Stack{
		element: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		len:     0,
	}
	// for i := len(arr) - 1; i >= 0; i-- {
	// 	s.push(arr[i])
	// }

	if !direction {
		for i := 0; i < len(arr); i++ {
			for {
				if s.isEmpty() {
					arrCopy[i] = 0
					s.push(i)
					break
				}
				if arr[i] <= arr[s.peek()] {
					s.pop()
				} else {
					arrCopy[i] = s.peek()
					s.push(i)
					break
				}
			}

		}
	} else {
		for i := len(arr) - 1; i >= 0; i-- {
			for {
				if s.isEmpty() {
					arrCopy[i] = len(arr) - 1
					s.push(i)
					break
				}
				if arr[i] <= arr[s.peek()] {
					s.pop()
				} else {
					arrCopy[i] = s.peek()
					s.push(i)
					break
				}
			}

		}
	}
	fmt.Println(arrCopy)
	return arrCopy
}
