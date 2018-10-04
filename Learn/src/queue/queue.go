package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}
