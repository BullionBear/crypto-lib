package hashlist

import (
	"github.com/BullionBear/crypto-lib/lib/cryptods/indexedlist"
	"github.com/emirpasic/gods/v2/lists/doublylinkedlist"
)

// Assert List implementation
var _ indexedlist.IndexedList[int, string] = (*HashList[int, string])(nil)

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
func (hl *HashList[K, V]) Put(key K, value V) {
	hl.store[key] = value
	hl.linkedList.Add(key)
}

// Get retrieves the value for the given key
func (hl *HashList[K, V]) Get(key K) (value V, found bool) {
	value, found = hl.store[key]
	return value, found
}

func (hl *HashList[K, V]) First() (K, V) {
	// Implementation will depend on the internal data structure used
	return 0, ""
}

// Remove deletes the value for the given key
func (hl *HashList[K, V]) Remove(key K) {
	delete(hl.store, key)
}

// AddBack adds the key-value pair to the end of the list
func (hl *HashList[K, V]) Append(key K, value V) {
	hl.store[key] = value
	hl.linkedList.Append(key)
}

// AddFront adds the key-value pair to the front of the list
func (hl *HashList[K, V]) Prepend(key K, value V) {
	hl.store[key] = value
	hl.linkedList.Prepend(key)
}

// RemoveBack removes the key-value pair from the end of the list
func (hl *HashList[K, V]) RemoveBack() {
	// Implementation will depend on the internal data structure used

	delete(hl.store, key)
	return
}

// RemoveFront removes the key-value pair from the front of the list
func (hl *HashList[K, V]) RemoveFront() {
	// Implementation will depend on the internal data structure used
	return
}
