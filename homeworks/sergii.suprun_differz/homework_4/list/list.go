package list

// List is the common interface implemented by LinkedList
type List interface {
	Add(e interface{})
	Insert(index int, e interface{})
	Set(index int, e interface{})
	Remove(e interface{}) bool
	Get(index int) interface{}
	Size() int
	IsEmpty() int
}
