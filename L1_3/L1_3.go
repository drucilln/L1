package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	sqrCh := make(chan int, len(arr))

	wg := sync.WaitGroup{}

	for _, val := range arr {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			sqrCh <- x * x
		}(val)
	}
	wg.Wait()
	close(sqrCh)

	sum := 0
	for ch := range sqrCh {
		sum += ch
	}
	fmt.Println(sum)
}
