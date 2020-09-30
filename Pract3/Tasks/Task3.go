package Tasks

import (
	"fmt"
	"math"
)

func findPrimes(max int) IntArray {
	size := max - 1;
	primes := IntArray{};
	erathosphen := make([]bool, size);

	for i := 0; i < size; i++ {
		if !erathosphen[i] {
			primes = append(primes, i + 2);
			for j := i; j < size; j += i + 2 {
				erathosphen[j] = true;
			}
		}
	}

	return primes;
}

func (array IntArray) contains(item int) bool {
	for _, v := range array {
		if v == item {
			return true;
		}
	}
	return false;
}

func (sample IntArray) from(array IntArray, prepare func (x int) int) IntArray {
	result := IntArray{};
	for _, v := range array {
		if sample.contains(prepare(v)) {
			result = append(result, v);
		}
	}
	return result;
}

func (arr IntArray) where(predicate func(x int) bool) IntArray {
	result := IntArray{};
	for _, v := range arr {
		if predicate(v) {
			result = append(result, v);
		}
	}
	return result;
}

func (sample IntArray) sort() IntArray {
	if len(sample) <= 1 {
		return sample;
	}

	less := sample.where(func(x int) bool {
		return x < sample[0];
	});

	equal := sample.where(func(x int) bool {
		return x == sample[0];
	})

	greater := sample.where(func(x int) bool {
		return x > sample[0];
	});

	return append(append(less.sort(), equal...), greater.sort()...);
}

func Task3() {
	array := randarr(10, -10, 10);
	fmt.Printf("Source array: %v\n", array);

	max := array.findBest(func (current int, found int) bool {
		return math.Abs(float64(found)) > math.Abs(float64(current));
	});
	max = int(math.Abs(float64(max)));

	allPrimes := findPrimes(max);

	primesInArray := allPrimes.from(array, func(x int) int {
		return int(math.Abs(float64(x)));
	});

	fmt.Printf("Primes from array: %v\n", primesInArray);

	sortedPrimesInArray := primesInArray.sort();
	fmt.Printf("Sorted primes from array: %v\n", sortedPrimesInArray);
}