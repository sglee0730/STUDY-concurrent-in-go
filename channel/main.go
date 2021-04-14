package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 6)
	go func() {
		for i := 0; i < 6; i++ {
			time.Sleep(time.Second * 1)
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
