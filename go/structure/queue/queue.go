package queue

type Queue interface {
	Put(v interface{})
	Pop() interface{}
	Size() int
}
