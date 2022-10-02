package algo

func merge(left, right []int) []int {
	temp := []int{}
	i := 0
	j := 0
	// fmt.Println(left, right)
	for {
		if i > len(left) || j > len(right) {
			break
		} else {
			if i < len(left) && j < len(right) && left[i] <= right[j] {
				temp = append(temp, left[i])
				i++
			} else if j < len(right) {
				temp = append(temp, right[j])
				j++
			} else {
				break
			}
		}
	}
	// fmt.Println(left, right)
	for {
		if i < len(left) {
			temp = append(temp, left[i])
			i++
		}
		if j < len(right) {
			temp = append(temp, right[j])
			j++
		}
		if i == len(left) && j == len(right) {
			return temp
		}
	}
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	start := 0
	end := len(arr)
	mid := (start + end) / 2
	left := MergeSort(arr[start:mid])
	right := MergeSort(arr[mid:end])
	// copy(arr, merge(left, right))

	return merge(left, right)

}
