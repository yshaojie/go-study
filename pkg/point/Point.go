package main

import "time"

func main() {
	//Range()

	tick := time.NewTicker(3 * time.Second)
	stopCh := make(chan struct{})
	go func() {
		time.Sleep(30 * time.Second)
		stopCh <- struct{}{}
	}()

	for {
	breakPoint:
		select {
		case dd := <-tick.C:
			println(dd.String())
		case <-stopCh:

			break breakPoint
			//default:
			//	println("default ....")
		}
	}

}

func Point() {
	age := 12
	println(&age)
	println(age)
	p := &age
	println(p)
	var pp = &age
	println(pp)
}

func Range() {
	nums := []int{2, 3, 4}
	for index, num := range nums {
		println("index=", index, " value=", num)
	}
}
