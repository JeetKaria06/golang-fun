package main

import (
	"fmt"
	"sync"
)

type SharedResource struct {
	mp   map[int]int
	mu   *sync.Mutex
	rwmu *sync.RWMutex
}

func run() {
	obj := SharedResource{
		mp:   make(map[int]int),
		mu:   &sync.Mutex{},
		rwmu: &sync.RWMutex{},
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go writeLoop(obj)
	// wg.Wait()
	go readLoop(obj)

	wg.Wait()
}

func readLoop(
	obj SharedResource,
) {
	for {
		obj.mu.Lock()
		for k, v := range obj.mp {
			fmt.Printf("Key: %d, Value: %d\n", k, v)
		}
		obj.mu.Unlock()
	}
}

func writeLoop(
	obj SharedResource,
) {
	for {
		obj.mu.Lock()
		for i := 0; i < 5; i += 1 {
			obj.mp[i] = i
		}
		obj.mu.Unlock()
	}
}
