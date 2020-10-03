package Tasks

import (
	"fmt"
	"strings"
)

type StringArray []string;

func readStrings(n int) StringArray {
	strings := make(StringArray, n);

	for i := 0; i < n; i++ {
		fmt.Print("Enter string: ");
		strings[i] = readText();
	}

	return strings;
}

func (strs StringArray) toRuneMatrix(width int) RuneMatrix {
	matrix := RuneMatrix(make([][]rune, len(strs)));
	for i, row := range strs {
		runes := []rune(row);
		if len(runes) > width {
			matrix[i] = runes[:width];
		} else if len(row) <= width {
			updated := row + strings.Repeat("*", width - len(runes));
			matrix[i] = []rune(updated);
		}
	}
	return matrix;
}

func (matrix RuneMatrix) toASCIIMatrix() [][]int8 {
	result := make([][]int8, len(matrix));

	for i, row := range matrix {
		result[i] = make([]int8, len(row));
		for j, char := range row {
			if char < 256 {
				result[i][j] = int8(char);
			}
		}
	}

	return result;
}

func print(matrix [][]int8) {
	for _, row := range matrix {
		for _, elem := range row {
			fmt.Printf("%d ", elem);
		}
		fmt.Print("\n");
	}
}

func Task6() {
	height := readIntNoError("Enter count of strings: ");
	strings := readStrings(height);

	width := readIntNoError("Enter maximum width: ");
	matrix := strings.toRuneMatrix(width);

	fmt.Print("Rune matrix: \n");
	matrix.print();

	asciiMatrix := matrix.toASCIIMatrix();
	fmt.Print("ASCII matrix: \n")
	print(asciiMatrix);
}
