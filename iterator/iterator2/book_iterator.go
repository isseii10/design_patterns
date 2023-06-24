package iterator2

type IBookIterator interface {
	HasNext() bool
	Next() *Book
}

type  BookIterator struct {
	books []*Book
	index int
}

func (b *BookIterator) HasNext() bool {
	return b.index < len(b.books)
}

// 現在の要素を返し、次の要素のindexを持つ
func (b *BookIterator) Next() *Book {
	if b.HasNext() {
		defer func() {
			b.index++
		}()
		return b.books[b.index]
	}
	return nil
}
