package main

import (
	"go-study/pkg/concurrent"
	"math/rand"
	"sync"
	"time"
)

func main() {
	locker := &sync.Mutex{}
	queue := concurrent.Queue{
		Capacity:     4,
		Items:        nil,
		ProductCond:  sync.NewCond(locker),
		ConsumerCond: sync.NewCond(locker),
	}
	wg := sync.WaitGroup{}
	count := 0
	mutex := sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			time.Sleep(time.Second * 3)
			defer wg.Done()
			for i := 0; i < 20000; i++ {
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
			for i := 0; i < 10000; i++ {
				queue.Push(rand.Int())
			}
		}()
	}

	wg.Wait()
	println(count)
}
