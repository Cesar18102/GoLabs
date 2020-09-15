package Tasks

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()));

func Task5() {
	arr := randarr(10, -5, 5);
	fmt.Printf("Randomized array: %v\n", arr);

	avg := avg(arr);
	fmt.Printf("AVG = %f\n", avg);

	var nonGreaterThanAvg = delete(arr, func(x float64) bool {
		return x > avg
	});
	fmt.Printf("Non-greater than AVG: %v\n", nonGreaterThanAvg);

	var pushed []float64 = nonGreaterThanAvg;
	pushedCount := int(math.Ceil(randval(3, 7)));
	fmt.Printf("Pushing %d elements\n", pushedCount);

	for i := 0; i < pushedCount; i++ {
		item := randval(-5, 5);
		pushed = pushTop(pushed, item);
		fmt.Printf("Pushed %f: %v\n", item, pushed);
	}

	shift := int(math.Ceil(randval(1, float64(len(pushed)))));
	shifted := shiftRight(pushed, shift);
	fmt.Printf("Shifted at %d: %v\n", shift, shifted);

	firstEven := first(shifted, func(x float64) bool {
		return int(x) % 2 == 0;
	});
	fmt.Printf("First even = %f\n", firstEven);
}

func first(arr []float64, predicate func(x float64) bool) float64 {
	for i := 0; i < len(arr); i++ {
		if predicate(arr[i]) {
			return arr[i];
		}
	}
	return -123456789;
}

func randval(min float64, max float64) float64 {
	return r.Float64() * (max - min) + min;
}

func randarr(count int, min float64, max float64) []float64 {
	arr := make([]float64, count);
	for i := 0; i < count; i++ {
		arr[i] = randval(min, max);
	}
	return arr;
}

func shiftRight(arr []float64, count int) []float64 {
	return append(arr[len(arr) - count:], arr[:len(arr) - count]...);
}

func pushTop(arr []float64, item float64) []float64 {
	result := make([]float64, len(arr) + 1);
	result[0] = item;
	copy(result[1:], arr);
	return result;
}

func delete(arr []float64, deletePredicate func(x float64) bool) []float64 {
	for i := 0; i < len(arr); i++ {
		if deletePredicate(arr[i]) {
			arr = append(arr[:i], arr[i+1:]...);
			i--;
		}
	}
	return arr;
}

func avg (arr []float64) float64 {
	sum := 0.0;
	for _, v := range arr {
		sum += v;
	}
	return sum / float64(len(arr));
}