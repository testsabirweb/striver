package array

import "fmt"

func SellBuy(arr []int) int {
	profit := 0
	buy := 0
	sell := 0
	i := 0
	tempProfit := 0
	// j := 0

	if len(arr) <= 1 {
		return -1
	}

	for i < len(arr)-1 {
		for i < len(arr)-1 {
			if arr[i] < arr[i+1] {
				buy = i
				break
			}
			i++
		}
		tempProfit = 0
		for i < len(arr)-1 {
			if arr[i] > arr[i+1] {
				sell = i
				break
			} else {
				tempProfit = tempProfit + (arr[i+1] - arr[i])
			}
			i++
		}
		profit = profit + tempProfit
		fmt.Println(buy, sell, profit, tempProfit)

	}

	return profit

}
