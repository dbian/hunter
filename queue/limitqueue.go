// FIFO queue
package queue

type LQ interface {
	Push(p data) bool
	Get(i int) data
}

type lq struct {
	arr []data
}

func NewLimitQueue() LQ {
	return &lq{}
}

var MaxLen = 20

func (l *lq) Push(p data) bool {
	for _, v := range l.arr {
		if p == v {
			return false
		}
	}
	l.arr = append(l.arr, p)
	if len(l.arr) == MaxLen+1 {
		l.arr = l.arr[1:]
		return true
	}

	return true
}
func (l *lq) Get(i int) data {
	return l.arr[i]
}
