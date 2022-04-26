package Tasks

import (
	"../NumIo"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()));

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

func makeNumFile(path string, generator func(i int) (float64, error), n int) *NumIo.NumFile {
	created, _ := os.Create(path);
	created.Close();

	file := NumIo.NewNumFile(path);

	defer func() {

	}()

	for i := 0; i < n; i++ {
		generated, err := generator(i);

		if(err != nil) {
			fmt.Printf("%f was not added due to %s\n", generated, err.Error());
			continue;
		}

		file.Add(generated);
		fmt.Printf("%f %s %s\n", generated, "was added to file", path);
	}

	file.Commit();

	return &file;
}

func Aggregate(arr []float64, init float64, aggregator func(acc float64, cur float64) float64) float64 {
	value := init;
	for _, item := range arr {
		value = aggregator(value, item);
	}
	return value;
}

func Best(arr []float64, better func(next float64, cur float64) bool) float64 {
	min := arr[0];
	for _, val := range arr {
		if(better(val, min)) {
			min = val;
		}
	}
	return min;
}

func Task11()  {
	file := makeNumFile(
		"C:/Users/Олег/Desktop/modules.txt",
		func(i int) (float64, error) {
			return float64(int(randval(0, 10))), nil;
		}, 10,
	);

	nums := file.Get();

	fmt.Printf("Nums in file: %f\n", nums);

	sumpow2 := Aggregate(nums, 0, func(acc float64, cur float64) float64 { return acc + cur * cur; });
	fmt.Printf("Sum of pows of 2 = %f\n", sumpow2);
}

func sqrt(x int) (float64, error) {
	if(x < 0) {
		return math.NaN(), errors.New("Arithmetic error");
	}
	return math.Sqrt(float64(x)), nil;
}

func Task12()  {
	file := makeNumFile(
		"C:/Users/Олег/Desktop/sqrts.txt",
		func(i int) (float64, error) {
			return sqrt(int(randval(-10, 10)));
		}, 10,
	);

	nums := file.Get();

	fmt.Printf("Nums in file: %f\n", nums);

	min := Best(nums, func(next float64, cur float64) bool { return next < cur; });
	max := Best(nums, func(next float64, cur float64) bool { return next > cur; });

	fmt.Printf("Sum of Min and Max = %f + %f = %f\n", min, max, min + max);
}

func div(x float64, y float64) (float64, error) {
	if(y == 0) {
		if(x == 0) {
			return math.NaN(), errors.New("Arithmetic error");
		}
		return math.Inf(1), errors.New("Division by zero");
	}
	return x / float64(y), nil;
}

func Task13()  {
	file := makeNumFile(
		"C:/Users/Олег/Desktop/pows.txt",
		func(i int) (float64, error) {
			k := int(randval(-10, 10));
			return div(math.Pow(-1, float64(k)) * math.Pow(2, float64(k)), float64(k));
		}, 10,
	);

	nums := file.Get();

	fmt.Printf("Nums in file: %f\n", nums);

	min := Best(nums, func(next float64, cur float64) bool { return next < cur; });
	max := Best(nums, func(next float64, cur float64) bool { return next > cur; });

	fmt.Printf("Sum of Min and Max = %f + %f = %f\n", min, max, min + max);
}

func Task1() {
	Task11();
	fmt.Print("\n****************\n\n");

	Task12();
	fmt.Print("\n****************\n\n");

	Task13();
	fmt.Print("\n****************\n\n");
}
