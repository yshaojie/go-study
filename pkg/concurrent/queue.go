package concurrent

import "sync"

type Queue struct {
	Capacity     int
	Items        []int
	ProductCond  *sync.Cond
	ConsumerCond *sync.Cond
}

func (q *Queue) Push(item int) {
	//q.ProductCond.L.Lock()
	//defer q.ProductCond.L.Unlock()
	//if len(q.Items) >= q.Capacity {
	//	q.ProductCond.Wait()
	//}
	q.Items = append(q.Items, item)
	q.ConsumerCond.Signal()
}

func (q *Queue) Pop() int {
	q.ConsumerCond.L.Lock()
	defer q.ConsumerCond.L.Unlock()
	println(len(q.Items))
	if len(q.Items) <= 0 {
		println("a")
		q.ConsumerCond.Wait()
	}
	item := q.Items[0]
	q.Items = q.Items[:1]
	//q.ProductCond.Signal()
	return item
}
