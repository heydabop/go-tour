package main

import(
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z -= (math.Pow(z, 2) - x)/(2*z)
	}
	return z
}

func main(){
	fmt.Println(sqrt(2))
}
