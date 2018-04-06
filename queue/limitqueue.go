// FIFO queue
package queue

type LQ interface {
	Push(p Data) bool
	Get(i int) Data
}

type lq struct {
	arr []Data
}

func NewLimitQueue() LQ {
	return &lq{}
}

var MaxLen = 20

func (l *lq) Push(p Data) bool {
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
func (l *lq) Get(i int) Data {
	return l.arr[i]
}
