package iterator1


type BookIterator struct {
	books []*Book
	index int
}

func (b *BookIterator) HasNext() bool {
	return b.index < len(b.books)
}

// 現在の要素を返し、次の要素のindexを持つ
func (b *BookIterator) Next() interface{} {
	if b.HasNext() {
		defer func() {
			b.index++
		}()
		return b.books[b.index]
	}
	return nil
}
