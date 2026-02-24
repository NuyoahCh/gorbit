package main

// MyArrayStack 用数组切片作为底层数据结构实现栈
type MyArrayStack[T any] struct {
	arr []T
}

// 向栈顶加入元素，时间复杂度 O(1)
func (s *MyArrayStack[T]) Push(e T) {
	s.arr = append(s.arr, e)
}

// 从栈顶弹出元素，时间复杂度 O(1)
func (s *MyArrayStack[T]) Pop() T {
	if len(s.arr) == 0 {
		var zero T
		return zero
	}
	e := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return e
}

// 查看栈顶元素，时间复杂度 O(1)
func (s *MyArrayStack[T]) Peek() T {
	if len(s.arr) == 0 {
		var zero T
		return zero
	}
	return s.arr[len(s.arr)-1]
}

// 返回栈中的元素个数，时间复杂度 O(1)
func (s *MyArrayStack[T]) Size() int {
	return len(s.arr)
}
