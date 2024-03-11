package queue

type Queue struct {
	list []any
}

func (q *Queue) Enq(item any) {
	q.list = append(q.list, item)
}

func (q *Queue) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue) Deq() (any, bool) {
	if !q.IsEmpty() {
		item := q.list[0]
		q.list = q.list[1:]
		return item, true
	}
	return nil, false
}
