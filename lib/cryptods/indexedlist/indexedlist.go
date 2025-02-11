/*
An indexed list is allows for a list of items to be stored in a map
and accessed by an index. This is useful for storing a list of items
that need to be accessed by an index, such as a list of items that
need to be accessed by a hash.
*/

package indexedlist

// IndexedList is an interface for a list of items that can be accessed by an index
type IndexedList[K comparable, V any] interface {
	Get(key K) (value V, found bool)
	Remove(key K)
	Append(key K, value V)
	Prepend(key K, value V)
	RemoveBack()
	RemoveFront()
}
