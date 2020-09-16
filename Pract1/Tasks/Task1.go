package Tasks

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()));

func randInt(min int, max int) int {
	return int(r.Float64() * float64(max - min) + float64(min));
}

func Task1()  {
	x, y, z := randInt(-10, 10), randInt(-10, 10), randInt(-10, 10);

	if anyOpposite(x, y, z) {
		fmt.Printf("There is a couple of opposite numbers in triplet of %d %d and %d", x, y, z);
	} else {
		fmt.Printf("There is no couple of opposite numbers in triplet of %d %d and %d", x, y, z);
	}
}

func anyOpposite(x int, y int, z int) bool {
	return opposite(x, y) || opposite(y, z) || opposite(x, z);
}

func opposite(x int, y int) bool {
	return x == -y;
}
