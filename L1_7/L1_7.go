package main

import (
	"fmt"
	"sync"
)

type Data struct {
	data map[int]int
	sync.Mutex
}

func NewData() *Data {
	return &Data{data: make(map[int]int, 100)}
}

func example1() {
	m := NewData()
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			m.Lock()
			m.data[x] = x
			m.Unlock()
		}(i)
	}
	wg.Wait()
	//m.Lock()
	//for k, v := range m.data {
	//	fmt.Println(k, v)
	//}
	//m.Unlock()
}

func example2() {
	m := make(map[int]int, 100)
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(x int, mu *sync.RWMutex) {
			defer wg.Done()
			mu.Lock()
			m[x] = x
			mu.Unlock()
		}(i, &mu)
	}
	wg.Wait()
}

func example3() {
	m := sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			m.Store(x, x)
		}(i)
	}
	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

func main() {
	example1()
	example2()
	example3()
}
