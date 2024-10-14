package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	chan1 := make(chan int)
	chan2 := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		defer close(chan1)
		for _, v := range arr {
			chan1 <- v
		}
	}()

	go func() {
		defer wg.Done()
		defer close(chan2)
		for v := range chan1 {
			chan2 <- v * 2
		}
	}()
	go func() {
		defer wg.Done()
		for v := range chan2 {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
