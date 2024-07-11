package hashlist

import (
	"github.com/BullionBear/crypto-lib/lib/cryptods/indexedlist"
	"github.com/emirpasic/gods/v2/lists/doublylinkedlist"
)

// Assert List implementation
var _ indexedlist.IndexedList[int, string] = (*HashList[int, string])(nil)

// HashList is a struct for storing a list of items that can be accessed by an index
type HashList[K comparable, V any] struct {
	linkedList *doublylinkedlist.List[K]
	store      map[K]V
}

// NewHashList creates a new HashList
func NewHashList[K comparable, V any]() *HashList[K, V] {
	return &HashList[K, V]{
		linkedList: doublylinkedlist.New[K](),
		store:      make(map[K]V),
	}
}

// Put inserts or updates the value for the given key
func (list *HashList[K, V]) Put(key K, value V) {
	list.store[key] = value
	list.linkedList.Add(key)
}

// Get retrieves the value for the given key
func (list *HashList[K, V]) Get(key K) (value V, found bool) {
	value, found = list.store[key]
	return value, found
}

// Remove deletes the value for the given key
func (list *HashList[K, V]) Remove(key K) {
	delete(list.store, key)
}

// AddBack adds the key-value pair to the end of the list
func (list *HashList[K, V]) Append(key K, value V) {
	list.store[key] = value
	list.linkedList.Append(key)
}

// AddFront adds the key-value pair to the front of the list
func (list *HashList[K, V]) Prepend(key K, value V) {
	list.store[key] = value
	list.linkedList.Prepend(key)
}

// RemoveBack removes the key-value pair from the end of the list
func (list *HashList[K, V]) RemoveBack() {
	var key K
	iter := list.linkedList.Iterator()
	if iter.Last() {
		key = iter.Value()
		delete(list.store, key)
	}
}

// RemoveFront removes the key-value pair from the front of the list
func (list *HashList[K, V]) RemoveFront() {
	var key K
	iter := list.linkedList.Iterator()
	if iter.First() {
		key = iter.Value()
		delete(list.store, key)
	}
}
