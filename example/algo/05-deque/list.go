package main

import (
	"container/list"
)

type MyListDeque struct {
	list *list.List
}

func NewMyListDeque() *MyListDeque {
	return &MyListDeque{list: list.New()}
}

// 从队头插入元素，时间复杂度 O(1)
func (d *MyListDeque) AddFirst(e interface{}) {
	d.list.PushFront(e)
}

// 从队尾插入元素，时间复杂度 O(1)
func (d *MyListDeque) AddLast(e interface{}) {
	d.list.PushBack(e)
}

// 从队头删除元素，时间复杂度 O(1)
func (d *MyListDeque) RemoveFirst() interface{} {
	if elem := d.list.Front(); elem != nil {
		return d.list.Remove(elem)
	}
	return nil
}

// 从队尾删除元素，时间复杂度 O(1)
func (d *MyListDeque) RemoveLast() interface{} {
	if elem := d.list.Back(); elem != nil {
		return d.list.Remove(elem)
	}
	return nil
}

// 查看队头元素，时间复杂度 O(1)
func (d *MyListDeque) PeekFirst() interface{} {
	if elem := d.list.Front(); elem != nil {
		return elem.Value
	}
	return nil
}

// 查看队尾元素，时间复杂度 O(1)
func (d *MyListDeque) PeekLast() interface{} {
	if elem := d.list.Back(); elem != nil {
		return elem.Value
	}
	return nil
}
