package main

import (
	"container/list"
)

// 用链表作为底层数据结构实现队列
type MyLinkedQueue struct {
	list *list.List
}

// 构造函数
func NewMyLinkedQueue() *MyLinkedQueue {
	return &MyLinkedQueue{list: list.New()}
}

// 向队尾插入元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Push(e interface{}) {
	q.list.PushBack(e)
}

// 从队头删除元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Pop() interface{} {
	front := q.list.Front()
	if front != nil {
		return q.list.Remove(front)
	}
	return nil
}

// 查看队头元素，时间复杂度 O(1)
func (q *MyLinkedQueue) Peek() interface{} {
	front := q.list.Front()
	if front != nil {
		return front.Value
	}
	return nil
}

// 返回队列中的元素个数，时间复杂度 O(1)
func (q *MyLinkedQueue) Size() int {
	return q.list.Len()
}
