package Tasks

import (
	"fmt"
	"math"
)

func Task6() {
	width := int(randval(2, 10));
	height := int(randval(2, 10));

	matrix := randmatrix(height, width, -5, 5);
	fmt.Printf("Randomized matrix: %v\n", matrix);

	rows := findRows(matrix, func(row []int) bool {
		negativeCount := 0;
		for _, v := range row {
			if(v < 0) {
				negativeCount++;
			}
		}
		return negativeCount % 3 != 0;
	});

	fmt.Print("Rows with negative count % 3 != 0\n");
	for _, row := range rows {
		elems := findElems(row, func(elem int) bool {
			return elem > 0 && elem % 2 != 0;
		});

		fmt.Printf("Row: %v; ", row);
		fmt.Printf("Odd elems > 0: %v\n", elems);
	}
}

func findElems(arr []int, predicate func(elem int) bool) []int {
	var result []int;
	for _, v := range arr {
		if predicate(v) {
			result = append(result, v);
		}
	}
	return result;
}

func findRows(arr [][]int, predicate func(row []int) bool) [][]int {
	var result [][]int;
	for _, row := range arr {
		if predicate(row) {
			result = append(result, row);
		}
	}
	return result;
}

func randmatrix(height int, width int, min int, max int) [][]int {
	arr := make([][]int, height);
	for i := 0; i < height; i++ {
		arr[i] = make([]int, width);
		for j := 0; j < width; j++ {
			arr[i][j] = int(math.Round(randval(float64(min), float64(max))));
		}
	}
	return arr;
}
