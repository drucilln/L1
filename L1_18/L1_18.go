package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterMutex struct {
	mu sync.Mutex
	x  int64
}

type CounterAtomic struct {
	x int64
}

func (ca *CounterAtomic) Inc() {
	atomic.AddInt64(&ca.x, 1)
}

func (ca *CounterAtomic) Value() int64 {
	return atomic.LoadInt64(&ca.x)
}

func (cm *CounterMutex) Inc() {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.x++
}

func (cm *CounterMutex) Value() int64 {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	return cm.x
}

func main() {
	var cm CounterMutex
	var ca CounterAtomic
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ca.Inc()
			cm.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(cm.Value())
	fmt.Println(ca.Value())
}
