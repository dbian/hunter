package queue

import (
	"testing"
)

func Test_lq_Insert(t *testing.T) {
	q := NewLimitQueue()
	q.Push(Data{Title: "a"})
	q.Push(Data{Title: "b"})
	q.Push(Data{Title: "c"})
	d := q.(*lq)
	if len(d.arr) != 3 {
		t.Error("insert fail", d.arr)
	}
	MaxLen = 3

	if !q.Push(Data{Title: "d"}) {
		t.Error("push fail")
	}
	if len(d.arr) != 3 {
		t.Error("insert fail", d.arr)
	}
	if q.Get(0).Title != "b" {
		t.Error("insert fail", d.arr)
	}
	if q.Get(1).Title != "c" {
		t.Error("insert fail", d.arr)
	}
	if q.Get(2).Title != "d" {
		t.Error("insert fail", d.arr)
	}

	if q.Push(Data{Title: "b"}) {
		t.Error("push should fail")
	}
}
