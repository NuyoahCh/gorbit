package main

type MyArrayQueue[E any] struct {
	arr *CycleArray[E]
}

func NewMyArrayQueue[E any]() *MyArrayQueue[E] {
	return &MyArrayQueue[E]{
		arr: NewCycleArray[E](),
	}
}

func (q *MyArrayQueue[E]) Push(t E) {
	q.arr.AddLast(t)
}

// func (q *MyArrayQueue[E]) Pop() E {
// 	return q.arr.RemoveFirst()
// }

// func (q *MyArrayQueue[E]) Peek() E {
// 	return q.arr.GetFirst()
// }

func (q *MyArrayQueue[E]) Size() int {
	return q.arr.Size()
}
