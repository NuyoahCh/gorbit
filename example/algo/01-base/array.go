package main

import (
	"errors"
	"fmt"
)

type MyArrayList struct {
	// data 存储元素的底层数组
	data []interface{}
	// size 当前元素的数量
	size int
}

// INIT_CAP 初始容量
const INIT_CAP = 1

// NewMyArrayList 创建一个新的 MyArrayList 实例，使用默认初始容量
func NewMyArrayList() *MyArrayList {
	return NewMyArrayListWithCapacity(INIT_CAP)
}

// NewMyArrayListWithCapacity 创建一个新的 MyArrayList 实例，使用指定的初始容量
func NewMyArrayListWithCapacity(initCapacity int) *MyArrayList {
	return &MyArrayList{
		data: make([]interface{}, initCapacity),
		size: 0,
	}
}

func (list *MyArrayList) AddLast(value interface{}) {
	cap := len(list.data)
	// 判断容量是否充足
	if list.size == cap {
		// TODO 进行二倍扩容
		list.resize(cap * 2)
	}
	// 将元素添加到末尾
	list.data[list.size] = value
	// 更新元素数量
	list.size++
}

func (list *MyArrayList) Add(index int, value interface{}) error {
	// TODO: 检查索引越界问题
	if err := list.checkPositionIndex(index); err != nil {
		return err
	}
	cap := len(list.data)
	// 判断容量是否充足
	if list.size == cap {
		// TODO 进行二倍扩容
		list.resize(cap * 2)
	}
	// 将元素添加到指定位置
	for i := list.size - 1; i >= index; i-- {
		list.data[i+1] = list.data[i]
	}
	// 将新元素插入到指定位置
	list.data[index] = value
	// 更新元素数量
	list.size++
	return nil
}

func (list *MyArrayList) AddFirst(value interface{}) error {
	return list.Add(0, value)
}

func (list *MyArrayList) RemoveLast() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("No such element")
	}
	cap := len(list.data)
	if list.size == cap/4 {
		// TODO 进行二倍缩容
		list.resize(cap / 2)
	}
	deleteVal := list.data[list.size-1]
	// 将末尾元素置空，帮助垃圾回收
	list.data[list.size-1] = nil
	// 更新元素数量
	list.size--
	return deleteVal, nil

}

func (list *MyArrayList) Remove(index int) (interface{}, error) {
	// TODO: 检查索引越界问题
	if err := list.checkPositionIndex(index); err != nil {
		return nil, err
	}
	cap := len(list.data)
	if list.size == cap/4 {
		// TODO 进行二倍缩容
		list.resize(cap / 2)
	}
	deleteVal := list.data[index]
	for i := index + 1; i < list.size; i++ {
		list.data[i-1] = list.data[i]
	}
	// 将末尾元素置空，帮助垃圾回收
	list.data[list.size-1] = nil
	// 更新元素数量
	list.size--
	return deleteVal, nil
}

func (list *MyArrayList) RemoveFirst() (interface{}, error) {
	return list.Remove(0)
}

func (list *MyArrayList) Get(index int) (interface{}, error) {
	// TODO: 检查索引越界问题
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	return list.data[index], nil
}

func (list *MyArrayList) Set(index int, value interface{}) (interface{}, error) {
	// TODO: 检查索引越界问题
	if err := list.checkElementIndex(index); err != nil {
		return nil, err
	}
	oldVal := list.data[index]
	list.data[index] = value
	return oldVal, nil
}

// 工具方法
func (list *MyArrayList) Size() int {
	return list.size
}

func (list *MyArrayList) IsEmpty() bool {
	return list.size == 0
}

// 将 data 的容量改为 newCap
func (list *MyArrayList) resize(newCap int) {
	temp := make([]interface{}, newCap)

	for i := 0; i < list.size; i++ {
		temp[i] = list.data[i]
	}

	list.data = temp
}

func (list *MyArrayList) isElementIndex(index int) bool {
	return index >= 0 && index < list.size
}

func (list *MyArrayList) isPositionIndex(index int) bool {
	return index >= 0 && index <= list.size
}

// 检查 index 索引位置是否可以存在元素
func (list *MyArrayList) checkElementIndex(index int) error {
	if !list.isElementIndex(index) {
		return fmt.Errorf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

// 检查 index 索引位置是否可以添加元素
func (list *MyArrayList) checkPositionIndex(index int) error {
	if !list.isPositionIndex(index) {
		return fmt.Errorf("Index: %d, Size: %d", index, list.size)
	}
	return nil
}

func (list *MyArrayList) Display() {
	fmt.Printf("size = %d cap = %d\n", list.size, len(list.data))
	fmt.Println(list.data)
}
