package main

import (
	"go-study/pkg/concurrent"
	"math/rand"
	"sync"
	"time"
)

func main() {
	queue := concurrent.Queue{
		Capacity:     100,
		Items:        nil,
		ProductCond:  sync.NewCond(&sync.Mutex{}),
		ConsumerCond: sync.NewCond(&sync.Mutex{}),
	}
	wg := sync.WaitGroup{}
	count := 0
	mutex := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			time.Sleep(time.Second * 3)
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				mutex.Lock()
				count = count + 1
				mutex.Unlock()
				queue.Pop()
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				queue.Push(rand.Int())
			}
		}()
	}

	wg.Wait()
	println(count)
}
