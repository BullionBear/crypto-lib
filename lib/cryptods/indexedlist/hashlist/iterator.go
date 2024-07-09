package hashlist

import (
	"github.com/emirpasic/gods/v2/containers"
	"github.com/emirpasic/gods/v2/lists/doublylinkedlist"
)

// Assert Iterator implementation
var _ containers.ReverseIteratorWithKey[string, int] = (*Iterator[string, int])(nil)

type Iterator[K comparable, V any] struct {
	iterator doublylinkedlist.Iterator[K]
	table    map[K]V
}
