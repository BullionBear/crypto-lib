package hashlist

import (
	"github.com/BullionBear/crypto-lib/lib/cryptods/indexedlist"
)

// Assert List implementation
var _ indexedlist.IndexedList[string, int] = (*HashList[string, int])(nil)

type HashList[K comparable, V any] struct {
	// You can define the fields of HashList here
	// Example: store map[K]V
	store map[K]V
}

// NewHashList creates a new HashList
func NewHashList[K comparable, V any]() *HashList[K, V] {
	return &HashList[K, V]{
		store: make(map[K]V),
	}
}

// Put inserts or updates the value for the given key
func (hl *HashList[K, V]) Put(key K, value V) {
	hl.store[key] = value
}

// Get retrieves the value for the given key
func (hl *HashList[K, V]) Get(key K) (value V, found bool) {
	value, found = hl.store[key]
	return
}

// Remove deletes the value for the given key
func (hl *HashList[K, V]) Remove(key K) {
	delete(hl.store, key)
}

// AddBack adds the key-value pair to the end of the list
func (hl *HashList[K, V]) AddBack(key K, value V) {
	// Implementation will depend on the internal data structure used
	hl.Put(key, value)
}

// AddFront adds the key-value pair to the front of the list
func (hl *HashList[K, V]) AddFront(key K, value V) {
	// Implementation will depend on the internal data structure used
	hl.Put(key, value)
}

// RemoveBack removes the key-value pair from the end of the list
func (hl *HashList[K, V]) RemoveBack() (key K, value V, found bool) {
	// Implementation will depend on the internal data structure used
	return
}

// RemoveFront removes the key-value pair from the front of the list
func (hl *HashList[K, V]) RemoveFront() (key K, value V, found bool) {
	// Implementation will depend on the internal data structure used
	return
}
