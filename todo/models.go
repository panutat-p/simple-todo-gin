package todo

import "gorm.io/gorm"

type Todo struct {
	Title string `json:"title"`
	gorm.Model
}

// TableName
// Gorm, provide custom table name
func (Todo) TableName() string {
	return "todolist"
}
