package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	timeSec := flag.Int("t", 1, "time in seconds")
	flag.Parse()

	wg := sync.WaitGroup{}

	dataChan := make(chan int)
	doneChan := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(*timeSec) * time.Second)
		close(doneChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sender(doneChan, dataChan)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		receiver(dataChan)
	}()
	wg.Wait()
}

func sender(doneChan <-chan struct{}, dataChan chan<- int) {
	i := 0
	for {
		select {
		case <-doneChan:
			close(dataChan)
			return
		case dataChan <- i:
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func receiver(dataChan <-chan int) {
	for v := range dataChan {
		fmt.Println(v)
	}
}
