package Tasks

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Task2()  {
	r := rand.New(rand.NewSource(time.Now().Unix()));
	a, b, c := r.Float64(), r.Float64(), r.Float64();

	fmt.Printf("triangle: %f, %f, %f\n", a, b, c);

	if isValidTriangle(a, b, c) {
		fmt.Printf("triangle is valid\n")
		fmt.Printf("triangle scope = %f", triangleScope(a, b, c));
	} else {
		fmt.Printf("triangle is invalid")
	}
}

func isValidTriangle(a float64, b float64, c float64) bool {
	return a + b > c && b + c > a && a + c > b;
}

func triangleScope(a float64, b float64, c float64) float64 {
	halfPerimeter := (a + b + c) / 2.0;
	return math.Sqrt(halfPerimeter * (halfPerimeter - a) * (halfPerimeter - b) * (halfPerimeter - c));
}
