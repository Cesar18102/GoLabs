package Tasks

import (
	"fmt"
	"math"
)

func printPoint(x interface{}, y interface{})  {
	if x == nil {
		x = 0;
	}

	if y == nil {
		y = 0;
	}

	fmt.Printf("Point (%d; %d)\n", x, y);
}

func avg(array ...int) float64 {
	sum := 0.0;
	for i := 0; i < len(array); i += 2 {
		sum += float64(array[i]);
	}
	count := (len(array) + 1) / 2.0;
	return float64(sum) / float64(count);
}

type FastSearchArray IntArray;

func (array FastSearchArray) find(item int) int {
	fmt.Print("fast ");
	if len(array) == 0 || (len(array) == 1 && array[0] != item) {
		return -1;
	}

	mid := len(array) / 2;

	if item > array[mid] {
		found := array[mid:].find(item);
		if found == -1 {
			return found;
		} else {
			return mid + found;
		}
	} else if item < array[mid] {
		return array[:mid].find(item);
	}

	return mid;
}

func (array IntArray) find(item int) int {
	fmt.Print("simple\n");
	for i, v := range array {
		if v == item {
			return i;
		}
	}
	return -1;
}

type Searchable interface {
	find(item int) int;
}

func solve(equation func(x float64) float64, start float64, end float64) float64 {
	mid := (end - start) / 2.0 + start;
	midres := equation(mid);

	if midres == 0.0 {
		return mid;
	}

	if midres > 0 != (equation(start) > 0) {
		return solve(equation, start, mid);
	}

	return solve(equation, mid, end);
}

func Task5() {
	printPoint(5, 3);
	printPoint(nil, nil);
	printPoint(5, nil);
	//</1>

	arr := randarr(10, -5, 5);
	fmt.Printf("Array: %v\n", arr);
	fmt.Printf("Average of even-located items = %f\n", avg(arr...));
	//</2>

	search := randval(-10, 10);

	arr2 := randarr(10, -10, 10).sort();
	fmt.Printf("Sorted array: %v\n", arr2);

	found := arr2.find(search);
	fmt.Printf("\n%d found at %d\n", search, found);

	arr2fast := FastSearchArray(arr2);
	fmt.Printf("Sorted array fast: %v\n", arr2fast);

	foundFast := arr2fast.find(search);
	fmt.Printf("%d fast found at %d\n", search, foundFast);
	//</3>

	var test Searchable;
	test = arr2;

	foundInterface := test.find(search);
	fmt.Printf("%d interface found at %d\n", search, foundInterface);
	//</4>

	equation := func (x float64) float64 {
		return x - 1 / (3 + math.Sin(3.6 * x));
	}
	solve := solve(equation, 0, 0.85);

	fmt.Printf("equation solve = %f\n", solve);
	//</5>
}
