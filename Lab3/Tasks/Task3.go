package Tasks

import (
	"fmt"
	"math"
)

var GOLD = (1 + math.Sqrt(5)) / 2;

func findMinGoldUnimodal(eq func (x float64) float64, eps float64, a float64, b float64, ch chan float64) {
	if(math.Abs(b - a) < eps) {
		ch <- (b + a) / 2;
		return;
	}

	x1, x2 := b - (b - a) / GOLD, a + (b - a) / GOLD;
	y1, y2 := eq(x1), eq(x2);

	if(y1 >= y2) {
		findMinGoldUnimodal(eq, eps, x1, b, ch);
		return;
	}

	findMinGoldUnimodal(eq, eps, a, x2, ch);
}

func findMinGold(eq func (x float64) float64, eps float64, a float64, b float64) []float64 {
	chans := make([]chan float64, 0);
	results := make([]float64, 0);
	last := a;


	for cur := a; cur <= b; cur += eps {
		wasGrowing := eq(last + eps) > eq(last);
		nowGrowing := eq(cur + eps) > eq(cur);

		if wasGrowing != nowGrowing {
			ch := make(chan float64);
			chans = append(chans, ch);
			go findMinGoldUnimodal(eq, eps, a, b, ch);

			last = cur;
		}
	}

	for _, ch := range chans {
		res := <- ch;
		results = append(results, res);
	}

	return results;
}

func Task3() {
	eq := func(x float64) float64 {
		return math.Abs(x * math.Log(x) - 1);
	}

	res := findMinGold(eq, 0.00001, 1.5, 2);
	fmt.Printf("results = %f", res);
}