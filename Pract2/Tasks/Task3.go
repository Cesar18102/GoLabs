package Tasks

import (
	"fmt"
)

func Task3() {
	matrix := randmatrix(5, 5, -5, 5);
	col := column(matrix, 3);
	maxcol := max(col);

	fmt.Printf("Max value of 3-rd column of matrix %v is %d", matrix, maxcol);
}

func max(arr []int) int {
	result := arr[0];
	for _, v := range arr {
		if v > result {
			result = v;
		}
	}
	return result;
}

func column(matrix [][]int, column int) []int {
	result := make([]int, len(matrix));
	for i, v := range matrix {
		result[i] = v[column];
	}
	return result;
}

func randmatrix(height int, width int, min int, max int) [][]int {
	matrix := make([][]int, height);
	for i := 0; i < height; i++ {
		matrix[i] = randarr(width, min, max);
	}
	return matrix;
}
