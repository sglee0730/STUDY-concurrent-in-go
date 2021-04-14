package main

import "fmt"

func main() {
	ch := generator(2,3)
	out := square(ch)

	for n := range out {
		fmt.Println(n)
	}
}

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n*n
		}
		close(out)
	}()
	return out
}