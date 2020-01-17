package hashmap

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestListNew(t *testing.T) {
	l := NewList()
	n := l.First()
	if n != nil {
		t.Error("First item of list should be nil.")
	}

	n = l.head.Next()
	if n != nil {
		t.Error("Next element of empty list should be nil.")
	}
}

func TestList_insertAt(t *testing.T) {
	//
	for i := 0; i < 200; i++ {
		el1 := &ListElement{
			key:     111,
			keyHash: 111,
		}
		el2 := &ListElement{
			key:     222,
			keyHash: 222,
		}
		el3 := &ListElement{
			key:     333,
			keyHash: 333,
		}
		newIl := &ListElement{
			key:     223,
			keyHash: 223,
		}
		l := NewList()
		l.Add(el1, nil)
		l.Add(el2, nil)
		l.Add(el3, nil)
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			rand.Seed(int64(time.Now().Nanosecond()))
			time.Sleep(time.Duration(rand.Intn(10)))
			l.Delete(el2)
			wg.Done()
		}()
		go func() {
			defer wg.Done()
			rand.Seed(int64(time.Now().Nanosecond()))
			time.Sleep(time.Duration(rand.Intn(10)))
			for {
				if _, inserted := l.Add(newIl, nil); inserted {
					return
				}
			}
		}()
		wg.Wait()

		if le := l.Len(); le != 3 {
			t.Error("l.Len() != 3", le)
			return
		}
		if _, found, _ := l.search(nil, newIl); found == nil {
			t.Error("newIl not found")
			return
		}
	}
}
