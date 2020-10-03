package Tasks

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin);

func readText() string {
	scanner.Scan();
	return scanner.Text();
}

func readInt() (int, error) {
	return strconv.Atoi(readText());
}

func readIntNoError(message string) int {
	fmt.Print(message);
	num, err := readInt();

	if err == nil {
		return num;
	} else {
		return readIntNoError(message);
	}
}

func readArray(readCountMsg string, readItemMsg string) []int {
	count := readIntNoError(readCountMsg);
	arr := make([]int, count);

	for i := 0; i < count; i++ {
		arr[i] = readIntNoError(readItemMsg);
	}

	return arr;
}

func insert(array []int, inserted int) []int {

	if inserted < array[0] {
		return append([]int{ inserted }, array...);
	} else if inserted >= array[len(array) - 1] {
		return append(array, inserted);
	}

	for i := 0; i < len(array); i++ {
		if array[i] > inserted {
			return append(append(append([]int{}, array[:i]...), inserted), array[i:]...);
		}
	}

	return array;
}

func Task1() {
	array := readArray("Enter items count: ", "Enter next item: ");
	sort.Ints(array);

	fmt.Printf("Entered array was sorted: %v\n", array);

	toInsert := readIntNoError("Enter a number to insert: ");
	array = insert(array, toInsert);

	fmt.Printf("%d inserted: %v", toInsert, array);
}