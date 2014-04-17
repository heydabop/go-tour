package main

import(
	"fmt"
	"net"
	"os"
	"time"
	"math/rand"
	"math"
	"math/cmplx"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var(
	ToBe bool      = false
	MaxInt uint64  = 1<<64 - 1
	z complex128   = cmplx.Sqrt(-5 + 12i)
)

const(
	Big    = 1 << 100
	Small  = Big >> 99
)

func needInt(x int) int {return x*10 + 1}
func needFloat(x float64) float64 {return x*0.1}

type Vertex struct {
	X int
	Y int
}

func main(){
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	rand.Seed(time.Now().Unix())
	fmt.Println("Welcome to the playground!")

	fmt.Println("This time is", time.Now())
	fmt.Println("Unix time is", time.Now().Unix())

	fmt.Println("And if you try to open a file:")
	fmt.Println(os.Open("hello.go"))

	fmt.Println("Or access the network:")
	fmt.Println(net.Dial("tcp", "google.com:80"))

	fmt.Println("Rand: ", rand.Intn(100))
	fmt.Println("Now you have %g problems.", math.Nextafter(2, 3))

	fmt.Println(math.Pi)

	fmt.Println(add(2, 3))

	fmt.Println(swap("hello", "world"))
	a, b := swap("this", "works")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Println(i, j, k, c, python, java)

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	//fmt.Println(needInt(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Small))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(Vertex{1, 2})
}
