package main

import "fmt"

func main(){
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 3 //overflow
	fmt.Println(<-c)
	fmt.Println(<-c)
}
