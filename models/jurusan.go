package models

import "time"

type Jurusan struct {
	ID            int
	Jurusan       string
	JurusanDetail JurusanDetail `gorm:"foreignKey:id;references:jurusan_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Jurusan) TableName() string {
	return "jurusan"
}
