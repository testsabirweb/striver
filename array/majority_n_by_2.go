package array

import "fmt"

func MajorityNBy2() {
	arr := []int{4, 4, 2, 4, 3, 4, 4, 3, 2, 4}
	element := 0
	count := 0
	for _, value := range arr {
		if element != value {
			if count == 0 {
				element = value
				count = 1
			} else {
				count--
			}
		} else {
			count++
		}
		if count >= len(arr)/2 {
			break
		}
	}
	fmt.Println(element)
}
