package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker1(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Горутина 1 завершила свою работу через runtime.Goexit")
	runtime.Goexit()
	for {
		fmt.Println("Горутина 1 работает ...")
		time.Sleep(1 * time.Second)
	}
}

func worker2(wg *sync.WaitGroup, doneChan <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-doneChan:
			fmt.Println("Горутина 2 завершила свою работу через done канал")
			return
		default:
			fmt.Println("Горутина 2 работает ...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker3(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина 3 завершила свою работу через Context.WithCancel")
			return
		default:
			fmt.Println("Горутина 3 работает ...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker4(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина 4 завершила свою работу через Context.WithTimeout")
			return
		default:
			fmt.Println("Горутина 4 работает ...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker5(wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("Горутина 5 завершила свою работу через таймер")
			return
		default:
			fmt.Println("Горутина 5 работает ...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	doneChan := make(chan struct{})
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	ctxWithTimeout, cancel1 := context.WithTimeout(ctx, time.Second*2)

	wg.Add(5)
	go worker1(&wg)

	go worker2(&wg, doneChan)
	go worker3(&wg, ctxWithCancel)
	go worker4(&wg, ctxWithTimeout)
	go worker5(&wg)
	time.Sleep(time.Second * 2)
	close(doneChan)
	cancel()
	cancel1()
	time.Sleep(time.Second * 1)

	wg.Wait()
}
