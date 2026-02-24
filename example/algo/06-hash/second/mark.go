package main

import (
	"errors"
	"fmt"
)

// KVNode 表示键值对的节点
type KVNode01 struct {
	key interface{}
	val interface{}
}

// MyLinearProbingHashMap2 是线性探查哈希表的结构体
type MyLinearProbingHashMap2 struct {
	table []*KVNode01
	size  int
}

// 被删除的 KVNode 的占位符
var dummy = &KVNode01{nil, nil}

// 默认的初始化容量
const initCap = 4

// 构造函数，初始化容量
func NewMyLinearProbingHashMap2(cap int) *MyLinearProbingHashMap2 {
	if cap <= 0 {
		cap = initCap
	}
	return &MyLinearProbingHashMap2{
		table: make([]*KVNode01, cap),
		size:  0,
	}
}

// **** 增/改 ****

// 添加 key -> val 键值对
// 如果键 key 已存在，则将值修改为 val
func (m *MyLinearProbingHashMap2) Put(key, val interface{}) error {
	if key == nil {
		return errors.New("key is null")
	}

	// 负载因子默认设为 0.75，超过则扩容
	if m.size >= len(m.table)*3/4 {
		m.resize(len(m.table) * 2)
	}

	index := m.getKeyIndex(key)
	if index != -1 {
		// key 已存在，修改对应的 val
		m.table[index].val = val
		return nil
	}

	// key 不存在
	x := &KVNode01{key, val}
	// 在 table 中找一个空位或者占位符，插入
	index = m.hash(key)
	for m.table[index] != nil && m.table[index] != dummy {
		index = (index + 1) % len(m.table)
	}
	m.table[index] = x
	m.size++
	return nil
}

// **** 删 ****

// 删除 key 和对应的 val，并返回 val
// 若 key 不存在，则返回 nil
func (m *MyLinearProbingHashMap2) Remove(key interface{}) error {
	if key == nil {
		return errors.New("key is null")
	}

	// 缩容
	if m.size < len(m.table)/8 {
		m.resize(len(m.table) / 2)
	}

	index := m.getKeyIndex(key)
	if index == -1 {
		// key 不存在，不需要 remove
		return nil
	}

	// 开始 remove
	// 直接用占位符表示删除
	m.table[index] = dummy
	m.size--
	return nil
}

// **** 查 ****

// 返回 key 对应的 val
// 如果 key 不存在，则返回 nil
func (m *MyLinearProbingHashMap2) Get(key interface{}) interface{} {
	if key == nil {
		return nil
	}

	index := m.getKeyIndex(key)
	if index == -1 {
		return nil
	}

	return m.table[index].val
}

// 检查是否包含指定的 key
func (m *MyLinearProbingHashMap2) ContainsKey(key interface{}) bool {
	return m.getKeyIndex(key) != -1
}

// 返回所有键的列表
func (m *MyLinearProbingHashMap2) Keys() []interface{} {
	keys := []interface{}{}
	for _, entry := range m.table {
		if entry != nil && entry != dummy {
			keys = append(keys, entry.key)
		}
	}
	return keys
}

// 返回哈希表中键值对的数量
func (m *MyLinearProbingHashMap2) Size() int {
	return m.size
}

// 对 key 进行线性探查，返回一个索引
// 根据 table[i] 是否为 nil 判断是否找到对应的 key
func (m *MyLinearProbingHashMap2) getKeyIndex(key interface{}) int {
	step := 0
	for i := m.hash(key); m.table[i] != nil; i = (i + 1) % len(m.table) {
		step++
		// 防止死循环
		if step > len(m.table) {
			// 这里可以触发一次 resize，把标记为删除的占位符清理掉
			m.resize(len(m.table))
			return -1
		}
		entry := m.table[i]
		// 遇到占位符直接跳过
		if entry == dummy {
			continue
		}
		if entry.key == key {
			return i
		}
	}
	return -1
}

// 哈希函数，将键映射到 table 的索引
// [0, len(table) - 1]
func (m *MyLinearProbingHashMap2) hash(key interface{}) int {
	return int(fmt.Sprintf("%v", key)[0]) % len(m.table)
}

// 扩容或缩容函数
func (m *MyLinearProbingHashMap2) resize(cap int) {
	newMap := NewMyLinearProbingHashMap2(cap)
	for _, entry := range m.table {
		if entry != nil && entry != dummy {
			newMap.Put(entry.key, entry.val)
		}
	}
	m.table = newMap.table
}
