package Tasks

import (
	"fmt"
	"strings"
)

type IRuneMatrix interface {
	getSize() (int, int);
	getItem(y int, x int) rune;
	setItem(y int, x int, elem rune);
	count(predicate func(matrix IRuneMatrix, y int, x int) bool) int;
	foreach(predicate func(matrix IRuneMatrix, y int, x int) bool, apply func(matrix IRuneMatrix, y int, x int));
}

type RuneMatrix [][]rune;

func (matrix RuneMatrix) getSize() (int, int) {
	if len(matrix) == 0 {
		return 0, 0;
	}

	return len(matrix), len(matrix[0]);
}

func (matrix RuneMatrix) getItem(y int, x int) rune {
	return matrix[y][x];
}

func (matrix RuneMatrix) setItem(y int, x int, elem rune) {
	matrix[y][x] = elem;
}

func (matrix RuneMatrix) count(predicate func(matrix IRuneMatrix, y int, x int) bool) int {
	count := 0;
	for i, row := range matrix {
		for j, _ := range row {
			if predicate(matrix, i, j) {
				count++
			}
		}
	}
	return count;
}

func (matrix RuneMatrix) foreach(predicate func(matrix IRuneMatrix, y int, x int) bool, apply func(matrix IRuneMatrix, y int, x int)) {
	for i, row := range matrix {
		for j, _ := range row {
			if predicate(matrix, i, j) {
				apply(matrix, i, j);
			}
		}
	}
}

func readRuneMatrix() RuneMatrix {
	matrix := RuneMatrix{};
	height, width := readIntNoError("Enter matrix height: "), -1;

	for i := 0; i < height; i++ {
		fmt.Print("Enter matrix row: ");
		row := []rune(readText());

		if width == -1 {
			width = len(row);
		}

		if len(row) > width {
			row = row[:width];
		} else if len(row) < width {
			tail := []rune(strings.Repeat("*", width - len(row)));
			row = append(row, tail...);
		}

		matrix = append(matrix, row);
	}

	return matrix;
}

func (matrix RuneMatrix) print() {
	for _, row := range matrix {
		fmt.Printf("%v\n", string(row));
	}
}

func Task4() {
	matrix := readRuneMatrix();

	aboveMainDiagonalPredicate := func(matrix IRuneMatrix, y int, x int) bool {
		return x > y;
	};

	underSideDiagonalPredicate := func(matrix IRuneMatrix, y int, x int) bool {
		height, _ := matrix.getSize();
		return x > height - y - 1;
	};

	aAboveMainDiagonalPredicate := func(matrix IRuneMatrix, y int, x int) bool {
		return aboveMainDiagonalPredicate(matrix, y, x) && matrix.getItem(y, x) == 'a';
	};

	bUnderSideDiagonalPredicate := func(matrix IRuneMatrix, y int, x int) bool {
		return underSideDiagonalPredicate(matrix, y, x) && matrix.getItem(y, x) == 'b';
	};

	aAboveMainCount, bUnderSideCount := matrix.count(aAboveMainDiagonalPredicate), matrix.count(bUnderSideDiagonalPredicate);

	fmt.Printf("count of 'a' symbols above the main diagonal is %d\n", aAboveMainCount);
	fmt.Printf("count of 'b' symbols under the side diagonal is %d\n", bUnderSideCount);

	matrix.foreach(aAboveMainDiagonalPredicate, func(matrix IRuneMatrix, y int, x int) {
		matrix.setItem(y, x, '*');
	});

	matrix.foreach(bUnderSideDiagonalPredicate, func(matrix IRuneMatrix, y int, x int) {
		matrix.setItem(y, x, '*');
	});

	fmt.Print("Modified matrix: \n");
	matrix.print();
}