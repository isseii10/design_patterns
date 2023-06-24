package main

import (
	"fmt"

	"github.com/isseii10/design_patterns/iterator/iterator2"
)

func main() {
	book1 := iterator2.NewBook("ボボボーボ・ボーボボ")
	book2 := iterator2.NewBook("バババーバ・バーババ")
	book3 := iterator2.NewBook("ビビビービ・ビービビ")

	books := iterator2.NewBooks()
	books.Append(book1)
	books.Append(book2)
	books.Append(book3)

	it := books.CreateIterator()
	for it.HasNext() {
		fmt.Println(it.Next().GetName())
	}
}