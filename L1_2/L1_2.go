package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x * x)

		}(arr[i])
	}

	wg.Wait()
}
