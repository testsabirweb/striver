package stackandqueue

import "fmt"

func NextGreaterElement() {
	var arr = []int{5, 7, 1, 7, 6, 0}
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	s := Stack{
		element: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		len:     0,
	}
	for i := len(arr) - 1; i >= 0; i-- {
		s.push(arr[i])
	}
	for i := len(arr) - 1; i >= 0; i-- {
		for {
			if s.isEmpty() {
				arrCopy[i] = -1
				s.push(arr[i])
				break
			}
			if arr[i] >= s.peek() {
				s.pop()
			} else {
				arrCopy[i] = s.peek()
				s.push(arr[i])
				break
			}
		}

	}
	fmt.Println(arrCopy)
}
