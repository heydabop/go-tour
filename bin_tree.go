package main

import "code.google.com/p/go-tour/tree"
import "fmt"

func walker(t *tree.Tree, ch chan int){
	if t.Left != nil {
		walker(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walker(t.Right, ch)
	}
}

// walk the tree t, sending all values to channel ch
func Walk(t *tree.Tree, ch chan int){
	walker(t, ch)
	close(ch)
}

// return true if t1 and t2 hold the same values
func Same(t1, t2*tree.Tree) bool{
	same := true
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		j, ok2 := <-ch2
		if !ok2 {
			same = false
			break
		}
		if i != j {
			same = false
		}
	}
	if _, ok2 := <-ch2; ok2 {
		same = false
	}
	return same
}

func main(){
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
