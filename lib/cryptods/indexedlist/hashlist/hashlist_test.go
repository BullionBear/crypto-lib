package hashlist

import (
	"testing"
)

func TestHashListNew(t *testing.T) {
	list := NewHashList[int, string]()

	if actualValue := len(list.store); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	if actualValue := list.linkedList.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestHashListPutAndGet(t *testing.T) {
	list := NewHashList[int, string]()
	list.Put(1, "one")
	list.Put(2, "two")
	list.Put(3, "three")

	if value, found := list.Get(1); !found || value != "one" {
		t.Errorf("Got %v expected %v", value, "one")
	}
	if value, found := list.Get(2); !found || value != "two" {
		t.Errorf("Got %v expected %v", value, "two")
	}
	if value, found := list.Get(4); found || value != "" {
		t.Errorf("Got %v expected %v", value, "")
	}
}

func TestHashListRemove(t *testing.T) {
	list := NewHashList[int, string]()
	list.Put(1, "one")
	list.Put(2, "two")
	list.Put(3, "three")
	list.Remove(2)

	if value, found := list.Get(2); found || value != "" {
		t.Errorf("Got %v expected %v", value, "")
	}
}

func TestHashListAppend(t *testing.T) {
	list := NewHashList[int, string]()
	list.Append(4, "four")

	if value, found := list.Get(4); !found || value != "four" {
		t.Errorf("Got %v expected %v", value, "four")
	}
}

func TestHashListPrepend(t *testing.T) {
	list := NewHashList[int, string]()
	list.Prepend(5, "five")

	if value, found := list.Get(5); !found || value != "five" {
		t.Errorf("Got %v expected %v", value, "five")
	}
}

func TestHashListRemoveBack(t *testing.T) {
	list := NewHashList[int, string]()
	list.Append(6, "six")
	list.RemoveBack()

	if value, found := list.Get(6); found || value != "" {
		t.Errorf("Got %v expected %v", value, "")
	}
}

func TestHashListRemoveFront(t *testing.T) {
	list := NewHashList[int, string]()
	list.Prepend(7, "seven")
	list.RemoveFront()

	if value, found := list.Get(7); found || value != "" {
		t.Errorf("Got %v expected %v", value, "")
	}
}

func TestIteratorNextAndPrev(t *testing.T) {
	list := NewHashList[int, string]()
	list.Put(1, "one")
	list.Put(2, "two")
	list.Put(3, "three")

	iter := list.Iterator()

	if !iter.First() {
		t.Errorf("Expected iterator to move to the first element")
	}

	if key, value := iter.Key(), iter.Value(); key != 1 || value != "one" {
		t.Errorf("Got key %v and value %v expected key %v and value %v", key, value, 1, "one")
	}

	if !iter.Next() {
		t.Errorf("Expected iterator to move to the next element")
	}

	if key, value := iter.Key(), iter.Value(); key != 2 || value != "two" {
		t.Errorf("Got key %v and value %v expected key %v and value %v", key, value, 2, "two")
	}

	if !iter.Prev() {
		t.Errorf("Expected iterator to move to the previous element")
	}

	if key, value := iter.Key(), iter.Value(); key != 1 || value != "one" {
		t.Errorf("Got key %v and value %v expected key %v and value %v", key, value, 1, "one")
	}
}

func TestIteratorNextTo(t *testing.T) {
	list := NewHashList[int, string]()
	list.Put(1, "one")
	list.Put(2, "two")
	list.Put(3, "three")

	iter := list.Iterator()
	iter.Begin()

	found := iter.NextTo(func(key int, value string) bool {
		return value == "two"
	})

	if !found {
		t.Errorf("Expected to find the element with value 'two'")
	}

	if key, value := iter.Key(), iter.Value(); key != 2 || value != "two" {
		t.Errorf("Got key %v and value %v expected key %v and value %v", key, value, 2, "two")
	}
}

func TestIteratorPrevTo(t *testing.T) {
	list := NewHashList[int, string]()
	list.Put(1, "one")
	list.Put(2, "two")
	list.Put(3, "three")

	iter := list.Iterator()
	iter.End()

	found := iter.PrevTo(func(key int, value string) bool {
		return value == "two"
	})

	if !found {
		t.Errorf("Expected to find the element with value 'two'")
	}

	if key, value := iter.Key(), iter.Value(); key != 2 || value != "two" {
		t.Errorf("Got key %v and value %v expected key %v and value %v", key, value, 2, "two")
	}
}

func TestIteratorBeginAndEnd(t *testing.T) {
	list := NewHashList[int, string]()
	list.Put(1, "one")
	list.Put(2, "two")
	list.Put(3, "three")

	iter := list.Iterator()

	iter.End()
	if iter.Prev() {
		key := iter.Key()
		value := iter.Value()
		if key != 3 || value != "three" {
			t.Errorf("Got key %v and value %v expected key %v and value %v", key, value, 3, "three")
		}
	}

	iter.Begin()
	if iter.Next() {
		key := iter.Key()
		value := iter.Value()
		if key != 1 || value != "one" {
			t.Errorf("Got key %v and value %v expected key %v and value %v", key, value, 1, "one")
		}
	}
}
