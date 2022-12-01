package models

type Learn struct {
	ID   string
	Name string
}

func (Learn) TableName() string {
	return "users"
}
