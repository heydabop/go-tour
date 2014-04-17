package main

import(
	"fmt"
	"math/cmplx"
)

func cbrt(x complex128) complex128 {
	z := x
	for i := 0; i < 10; i++ {
		z -= (cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))
	}
	return z;
}

func main(){
	fmt.Println(cbrt(2))
	fmt.Println(cmplx.Pow(cbrt(2), 3))
}
