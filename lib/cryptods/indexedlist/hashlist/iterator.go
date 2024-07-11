package hashlist

import (
	"github.com/emirpasic/gods/v2/containers"
	"github.com/emirpasic/gods/v2/lists/doublylinkedlist"
)

// Assert Iterator implementation
var _ containers.ReverseIteratorWithKey[string, int] = (*Iterator[string, int])(nil)

type Iterator[K comparable, V any] struct {
	list     *HashList[K, V]
	iterator doublylinkedlist.Iterator[K]
}

func (list *HashList[K, V]) Iterator() Iterator[K, V] {
	return Iterator[K, V]{
		list:     list,
		iterator: list.linkedList.Iterator(),
	}
}

func (iter *Iterator[K, V]) Next() bool {
	return iter.iterator.Next()
}

func (iter *Iterator[K, V]) Prev() bool {
	return iter.iterator.Prev()
}

func (iter *Iterator[K, V]) Begin() {
	iter.iterator.Begin()
}

func (iter *Iterator[K, V]) End() {
	iter.iterator.End()
}

func (iter *Iterator[K, V]) First() bool {
	return iter.iterator.First()
}

func (iter *Iterator[K, V]) Last() bool {
	return iter.iterator.Last()
}

func (iter *Iterator[K, V]) Key() K {
	return iter.iterator.Value()
}

func (iter *Iterator[K, V]) Value() V {
	key := iter.iterator.Value()
	return iter.list.store[key]
}

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
