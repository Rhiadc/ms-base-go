package sql

import "github.com/google/uuid"

type Book struct {
	ID     uuid.UUID
	Title  string
	Pages  string
	Author string
}
