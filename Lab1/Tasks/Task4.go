package Tasks

import (
	"fmt"
	"math"
)

type Point struct {
	X float64;
	Y float64;
}

func Task4() {
	const START = 0.5;
	const END = 3.1;
	const DELTA = 0.21;

	const START_PARAM = -1.4;
	const END_PARAM = -0.1;
	const DELTA_PARAM = 0.1;

	const B = 1.5;

	for a := START_PARAM; a <= END_PARAM; a += DELTA_PARAM {
		var function func(float64) float64 = func(x float64) float64 {
			if x < B {
				return math.Sin(math.Abs(a * x + math.Pow(B, a)));
			} else {
				return math.Cos(math.Abs(a * x - math.Pow(B, a)));
			}
		};

		points := calculatePoints(function, START, END, DELTA);
		fmt.Printf("a = %f, points: %v\n", a, points);
	}
}

func calculatePoints(function func(float64) float64, start float64, end float64, delta float64) []Point {
	result := []Point { };

	for i := start; i <= end; i += delta {
		coord := Point { i, function(i) };
		result = append(result, coord);
	}

	return result;
}
