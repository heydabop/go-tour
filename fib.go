package main

import "fmt"

func fibonacci() func() int {
	var last, curr int = 0, 1
	return func() int {
		temp := curr
		curr = last + curr
		last = temp
		return curr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
