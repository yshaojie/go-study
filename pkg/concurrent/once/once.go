package main

import "sync"

func main() {
	// 保证代码只执行一次
	once := sync.Once{}
	for i := 0; i < 10000; i++ {
		go func() {
			once.Do(func() {
				println("aaaaaaa")
			})
		}()
	}
	//map1 := sync.Pool{}

}
