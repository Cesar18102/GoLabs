package Tasks

import (
	"fmt"
	"math"
)

type FloatArray []float64;

func determine(function func (x float64) float64, start float64, end float64, delta float64) []float64 {
	result := []float64{};
	for i := start; i < end; i += delta {
		result = append(result, function(i));
	}
	return result;
}

func (arr FloatArray) where(predicate func(x float64) bool) FloatArray {
	result := FloatArray{};
	for _, v := range arr {
		if predicate(v) {
			result = append(result, v);
		}
	}
	return  result;
}

func (arr FloatArray) combine(other FloatArray, combiner func (x float64, y float64) float64) FloatArray {
	combined := FloatArray{};
	for _, v := range arr {
		for _, vother := range other {
			combined = append(combined, combiner(v, vother));
		}
	}
	return combined;
}

func (arr FloatArray) findBest(better func(current float64, found float64) bool) float64 {
	best := arr[0];
	for i := 1; i < len(arr); i++ {
		if better(best, arr[i]) {
			best = arr[i];
		}
	}
	return best;
}

//Сформировать массив из суммы отрицательных элементов массивов X, Y, W. Найти максимальную сумму.
func Task2() {
	x := determine(func (x float64) float64 {
		return 3 * x * x - 9 * (5 - x);
	}, 0, 8, 1);
	fmt.Printf("X = %f\n", x);

	y := determine(func (x float64) float64 {
		return 6 * math.Sin(2 * x) + 5 * math.Pow(math.E, x - 5);
	}, 0, 8, 1);
	fmt.Printf("Y = %f\n", y);

	z := determine(func (x float64) float64 {
		return 6 * (x - 4) + 5 * math.Pow(math.Sin(1.5 * x), 2.0);
	}, 0, 10, 1);
	fmt.Printf("Z = %f\n", z);

	w := determine(func (x float64) float64 {
		return 6 * math.Log(0.1 * x) + 10 * math.Pow(x - 2.5, 3.0);
	}, 0, 15, 1);
	fmt.Printf("W = %f\n", w);

	negativePredicate := func(x float64) bool {
		return x < 0;
	}

	xn, yn, zn, wn := FloatArray(x).where(negativePredicate), FloatArray(y).where(negativePredicate), FloatArray(z).where(negativePredicate), FloatArray(w).where(negativePredicate);

	fmt.Printf("Negative from X: %f\n", xn);
	fmt.Printf("Negative from Y: %f\n", yn);
	fmt.Printf("Negative from Z: %f\n", zn);
	fmt.Printf("Negative from W: %f\n", wn);

	summator := func(x float64, y float64) float64 {
		return x + y;
	}
	xywn := xn.combine(yn, summator).combine(wn, summator);
	fmt.Printf("Sums of negatives: %f\n", xywn);

	largestSum := xywn.findBest(func(current float64, found float64) bool {
		return found > current;
	});
	fmt.Printf("The largest sum is %f", largestSum);
}