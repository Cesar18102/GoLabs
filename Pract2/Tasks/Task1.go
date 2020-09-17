package Tasks

import (
	"fmt"
	"math/rand"
	"time"
)

func Task1() {
	arr := randarr(10, -5, 5);
	even := where(arr, func(x int) bool{
		return x % 2 == 0;
	});
	result := mult(even);

	fmt.Printf("Event elems of %v are %v; multiplication = %d", arr, even, result);
}

func where(arr []int, predicate func(x int) bool) []int {
	found := make([]int, 0);
	for _, v := range arr {
		if predicate(v) {
			found = append(found, v);
		}
	}
	return found;
}

func mult(arr []int) int {
	result := 1;
	for _, v := range arr {
		result *= v;
	}
	return result;
}

var r = rand.New(rand.NewSource(time.Now().Unix()));

func randval(min int, max int) int {
	return int(r.Float64() * float64(max - min) + float64(min));
}

func randarr(n int, min int, max int) []int {
	arr := make([]int, n);
	for i := 0; i < n; i++ {
		arr[i] = randval(min, max);
	}
	return arr;
}
