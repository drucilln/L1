package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func worker(n int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d received job %d\n", n, job)
	}
}

func main() {
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}

	jobs := make(chan int)

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	wg.Add(1)
	go func() {
		defer wg.Done()
		genJobs(sigChan, jobs)
	}()
	wg.Wait()
	fmt.Println("All worker`s done")
}

func genJobs(sigChan chan os.Signal, jobs chan<- int) {
	defer close(jobs)
	i := 0
	for {
		select {
		case <-sigChan:
			fmt.Println("Received terminate signal")
			return
		default:
			time.Sleep(100 * time.Millisecond)
			jobs <- i
			i++
		}
	}
}
