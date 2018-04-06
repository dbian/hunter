package queue

import (
	"testing"
)

func Test_lq_Insert(t *testing.T) {
	q := NewLimitQueue()
	q.Push(data{title: "a"})
	q.Push(data{title: "b"})
	q.Push(data{title: "c"})
	d := q.(*lq)
	if len(d.arr) != 3 {
		t.Error("insert fail", d.arr)
	}
	MaxLen = 3

	if !q.Push(data{title: "d"}) {
		t.Error("push fail")
	}
	if len(d.arr) != 3 {
		t.Error("insert fail", d.arr)
	}
	if q.Get(0).title != "b" {
		t.Error("insert fail", d.arr)
	}
	if q.Get(1).title != "c" {
		t.Error("insert fail", d.arr)
	}
	if q.Get(2).title != "d" {
		t.Error("insert fail", d.arr)
	}

	if q.Push(data{title: "b"}) {
		t.Error("push should fail")
	}
}
