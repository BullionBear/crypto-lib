/*
An indexed list is allows for a list of items to be stored in a map
and accessed by an index. This is useful for storing a list of items
that need to be accessed by an index, such as a list of items that
need to be accessed by a hash.
*/

package indexedlist

type IndexedList[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	AddBack(key K, value V)
	AddFront(key K, value V)
	RemoveBack() (key K, value V, found bool)
	RemoveFront() (key K, value V, found bool)
}
