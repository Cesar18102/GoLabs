//Tasks is a package
package Tasks

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Task1()  {
	r := rand.New(rand.NewSource(time.Now().Unix()));

	length := r.Float64();
	scope := circleScopeByLength(length);

	fmt.Printf("length = %f, scope = %f", length, scope);
}

func circleScopeByLength(length float64) float64 {
	radius := length / (math.Pi * 2.0);
	return math.Pi * math.Pow(radius, 2.0);
}