package Tasks

import (
	"fmt"
)

func (matrix IntMatrix) apply(function func(row IntArray) IntArray, predicate func(rownum int, row IntArray) bool) {
	for i, row := range matrix {
		if predicate(i, row) {
			matrix[i] = function(row);
		}
	}
}
func (matrix IntMatrix) print() {
	for _, row := range matrix {
		fmt.Printf("%d\n", row);
	}
	fmt.Print("\n");
}

func Task4() {
	shift := randval(1, 5);
	matrix := randmatrix(5, 5, -5, 5);

	fmt.Print("Matrix: \n");
	matrix.print();

	matrix.apply(func(row IntArray) IntArray {
		return append(row[len(row) - shift:], row[:len(row) - shift]...);
	}, func(rownum int, row IntArray) bool {
		return rownum % 2 == 0;
	});

	fmt.Printf("Matrix with even row shifted by %d: \n",shift);
	matrix.print();
}
