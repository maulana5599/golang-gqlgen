package models

type BookDetail struct {
	ID       int
	BookId   int
	BookDesc string
}

func (BookDetail) TableName() string {
	return "book_detail"
}
