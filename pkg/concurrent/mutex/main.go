package main

import "sync"

func main() {
	threadSafe()
	println("\n\n\n\n")
	theadNotSafe()
}

func theadNotSafe() {
	for i := 0; i < 100; i++ {
		count := 0
		wg := sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 10000; i++ {
					count = count + 1
				}
				wg.Done()
			}()
		}
		wg.Wait()
		println("theadNotSafe count=", count)
	}

}

func threadSafe() {
	mutex := sync.Mutex{}
	for i := 0; i < 100; i++ {
		count := 0
		wg := sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 10000; i++ {
					mutex.Lock()
					count = count + 1
					mutex.Unlock()
				}
				wg.Done()
			}()
		}
		wg.Wait()
		println("threadSafe count=", count)
	}

}
