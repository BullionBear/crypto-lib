package hashlist

import (
	"github.com/emirpasic/gods/v2/containers"
	"github.com/emirpasic/gods/v2/lists/doublylinkedlist"
)

// Assert Iterator implementation
var _ containers.ReverseIteratorWithKey[string, int] = (*Iterator[string, int])(nil)

// Iterator is a struct for iterating over the HashList
func (list *HashList[K, V]) Iterator() Iterator[K, V] {
	return Iterator[K, V]{
		list:     list,
		iterator: list.linkedList.Iterator(),
	}
}

// Iterator is a struct for iterating over the HashList
type Iterator[K comparable, V any] struct {
	list     *HashList[K, V]
	iterator doublylinkedlist.Iterator[K]
}

// Next moves to the next element
func (list *Iterator[K, V]) Next() bool {
	return list.iterator.Next()
}

// Prev moves to the previous element
func (list *Iterator[K, V]) Prev() bool {
	return list.iterator.Prev()
}

// Begin moves to before the first element
func (list *Iterator[K, V]) Begin() {
	list.iterator.Begin()
}

// End moves to after the last element
func (list *Iterator[K, V]) End() {
	list.iterator.End()
}

// First moves to the first element
func (list *Iterator[K, V]) First() bool {
	return list.iterator.First()
}

// Last moves to the last element
func (list *Iterator[K, V]) Last() bool {
	return list.iterator.Last()
}

// Key returns the key of the current element
func (list *Iterator[K, V]) Key() K {
	return list.iterator.Value()
}

// Value returns the value of the current element
func (list *Iterator[K, V]) Value() V {
	key := list.iterator.Value()
	return list.list.store[key]
}

// PrevTo moves to the previous element that satisfies the condition
func (list *Iterator[K, V]) PrevTo(cond func(key K, value V) bool) bool {
	for list.Prev() {
		key := list.iterator.Value()
		value := list.list.store[key]
		if cond(key, value) {
			return true
		}
	}
	return false
}

// NextTo moves to the next element that satisfies the condition
func (list *Iterator[K, V]) NextTo(cond func(key K, value V) bool) bool {
	for list.Next() {
		key := list.iterator.Value()
		value := list.list.store[key]
		if cond(key, value) {
			return true
		}
	}
	return false
}
