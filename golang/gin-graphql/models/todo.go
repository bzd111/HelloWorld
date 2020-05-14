package models

import (
	"time"
)

type Todo struct {
	ID        int
	Text      string
	Done      bool
	UserID    int
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (Todo) TableName() string {
//     return "todos"
// }
