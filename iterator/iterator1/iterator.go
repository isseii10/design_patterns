package iterator1

type Iterator interface {
	HasNext() bool
	Next() interface{}
}