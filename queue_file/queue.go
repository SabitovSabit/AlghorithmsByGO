package queue_pck

type Queue struct{
	items []int
}

func (q *Queue) Enqueue(num int){
   q.items=append(q.items, num)
}

func (q *Queue) Dequeue() int{
	removeItem:=q.items[0]
	q.items=q.items[1:len(q.items)]
	return removeItem
}

func (q *Queue) ReturnAll() []int{
	return q.items
}