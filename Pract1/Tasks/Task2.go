package Tasks

import (
	"fmt"
	"math"
)

func Task2()  {
	n := randInt(100, 1000);
	sum := digitSum(n);

	fmt.Printf("Sum of digits of %d is %d (%v)", n, sum, numberParityText(sum));
}

func numberParityText(n int) string {
	if n % 2 == 0 {
		return "even";
	} else {
		return "odd";
	}
}

func digitSum(n int) int {
	sum, digits := 0, int(math.Ceil(math.Log10(float64(n))));
	for i := 0; i < digits; i++ {
		sum += n % 10;
		n /= 10;
	}
	return sum;
}