package Tasks

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type IntArray []int;
type IntMatrix []IntArray;

var r = rand.New(rand.NewSource(time.Now().Unix()));

func randval(min int, max int) int {
	return int(r.Float64() * float64(max - min) + float64(min));
}

func randarr(n int, min int, max int) IntArray {
	arr := make([]int, n);
	for i := 0; i < n; i++ {
		arr[i] = randval(min, max);
	}
	return arr;
}

func randmatrix(height int, width int, min int, max int) IntMatrix {
	matrix := make([]IntArray, height);
	for i := 0; i < height; i++ {
		matrix[i] = randarr(width, min, max);
	}
	return matrix;
}

func (arr IntArray) findBest(better func(current int, found int) bool) int {
	best := arr[0];
	for i := 1; i < len(arr); i++ {
		if better(best, arr[i]) {
			best = arr[i];
		}
	}
	return best;
}

func (matrix IntMatrix)findBest(better func(current int, found int) bool) int {
	best := matrix[0].findBest(better);
	for i := 1; i < len(matrix); i++ {
		found := matrix[i].findBest(better);
		if better(best, found) {
			best = found;
		}
	}
	return best;
}

func Task1() {
	a, b, c := randmatrix(5, 5, -10, 10), randmatrix(5, 5, -10, 10), randmatrix(5, 5, -10, 10);
	better := func(current int, found int) bool {
		return math.Abs(float64(found)) > math.Abs(float64(current));
	};

	fmt.Printf("Matrix A: %v\n", a);
	fmt.Printf("Matrix B: %v\n", b);
	fmt.Printf("Matrix C: %v\n", c);

	abest, bbest, cbest := a.findBest(better), b.findBest(better), c.findBest(better);

	fmt.Printf("||A|| + ||B|| + ||C|| = %d + %d + %d = %d\n", abest, bbest, cbest, abest + bbest + cbest);
}
