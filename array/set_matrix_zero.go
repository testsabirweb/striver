package array

import "fmt"

func SetMatrixZero() {
	a := [][]int{
		{1, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	row := make([]int, len(a))
	col := make([]int, len(a[0]))
	for index, _ := range row {
		row[index] = 1
	}

	for index, _ := range col {
		col[index] = 1
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == 0 {
				row[i] = 0
				col[j] = 0
			}
		}
	}

	fmt.Println(row, col)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if row[i] == 0 || col[j] == 0 {
				a[i][j] = 0
			}
		}

	}

	fmt.Println(a)
}
