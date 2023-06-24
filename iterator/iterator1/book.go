package iterator1


type Book struct {
	name string
}
func NewBook(name string) *Book {
	return &Book{name: name}
}
func (b Book) GetName() string {
	return b.name
}

// Bookのコレクション
type Books struct {
	bookList []*Book
}

func NewBooks() *Books {
	bookList := make([]*Book, 0)
	return &Books{
		bookList: bookList,
	}
}

func (b *Books) Append(book *Book) {
	b.bookList = append(b.bookList, book)
}

func (b *Books) CreateIterator() Iterator {
	return &BookIterator{
		books: b.bookList,
		index: 0,
	}
}
