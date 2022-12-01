package models

import "time"

type JurusanDetail struct {
	ID            int
	JurusanId     int
	JurusanDetail string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (JurusanDetail) TableName() string {
	return "jurusan_detail"
}
