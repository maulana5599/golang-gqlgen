package models

type Books struct {
	ID         int
	BookOwner  int
	BookName   string
	BookCover  string
	BookDetail BookDetail `gorm:"foreignKey:id;references:book_id"`

	// gorm:"foreignKey:berisi_local_id;references:berisi_fk_table_tujuan"
}
