package LinkedList

type LinkedList interface {
	First() (interface{}, error)
	Last() (interface{}, error)
	Length() int
	Prepend(value interface{})
	Append(value interface{})
	Remove(value interface{}) (interface{}, error)
	Contains(value interface{}) bool
}
