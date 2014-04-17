package main

import(
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error: %f does not have any real roots.", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	for i := 0; i < 10; i++ {
		z -= (math.Pow(z, 2) - x)/(2*z)
	}
	return z, nil
}

func main(){
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
