package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	count := 10
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			wg.Done()
		}()
	}
	println("wait for task complete", time.Now().String())
	wg.Wait()
	println("task complete .... ", time.Now().String())

}
