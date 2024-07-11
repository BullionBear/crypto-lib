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
func (iter *Iterator[K, V]) Next() bool {
	return iter.iterator.Next()
}

// Prev moves to the previous element
func (iter *Iterator[K, V]) Prev() bool {
	return iter.iterator.Prev()
}

// Begin moves to before the first element
func (iter *Iterator[K, V]) Begin() {
	iter.iterator.Begin()
}

// End moves to after the last element
func (iter *Iterator[K, V]) End() {
	iter.iterator.End()
}

// First moves to the first element
func (iter *Iterator[K, V]) First() bool {
	return iter.iterator.First()
}

// Last moves to the last element
func (iter *Iterator[K, V]) Last() bool {
	return iter.iterator.Last()
}

// Key returns the key of the current element
func (iter *Iterator[K, V]) Key() K {
	return iter.iterator.Value()
}

// Value returns the value of the current element
func (iter *Iterator[K, V]) Value() V {
	key := iter.iterator.Value()
	return iter.list.store[key]
}

// PrevTo moves to the previous element that satisfies the condition
func (iter *Iterator[K, V]) PrevTo(cond func(key K, value V) bool) bool {
	for iter.Prev() {
		key := iter.iterator.Value()
		value := iter.list.store[key]
		if cond(key, value) {
			return true
		}
	}
	return false
}

// NextTo moves to the next element that satisfies the condition
func (iter *Iterator[K, V]) NextTo(cond func(key K, value V) bool) bool {
	for iter.Next() {
		key := iter.iterator.Value()
		value := iter.list.store[key]
		if cond(key, value) {
			return true
		}
	}
	return false
}
