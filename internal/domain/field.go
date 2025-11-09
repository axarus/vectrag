package domain

import "time"

type Field struct {
	ID          string
	Name        string
	Type        FieldType
	Description string
	Unique      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
