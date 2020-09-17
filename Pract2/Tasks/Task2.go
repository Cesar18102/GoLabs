package Tasks

import "fmt"

func Task2() {
	arr := randarr(18, -10, 10);
	negative := whereslice(arr, func(x int) bool {
		return x < 0;
	});

	fmt.Printf("Negative elements of %v are %v", arr, negative);
}

func whereslice(arr []int, predicate func(x int) bool) []int {
	result := make([]int, 0);
	for i := 0; i < len(arr); i++ {
		if predicate(arr[i]) {
			result = append(result, arr[i:i + 1]...);
		}
	}
	return  result;
}
