package sql

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBookRepository(t *testing.T) {
	db, err := sql.Open("postgres", "postgresql://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	repo := NewBookRepository(db)
	assert.NotNil(t, repo, "Expected NewBookRepository to return a non-nil instance")
	assert.Equal(t, db, repo.db, "Expected db to be the same as the one provided")
}

func TestNewBookRepositoryWithNilDb(t *testing.T) {
	repo := NewBookRepository(nil)
	assert.NotNil(t, repo, "Expected NewBookRepository to return a non-nil instance even with nil db")
}

func TestNewBookRepositoryDbIntegrity(t *testing.T) {
	db, err := sql.Open("postgres", "postgresql://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	originalDbPointer := db
	repo := NewBookRepository(db)
	assert.Equal(t, originalDbPointer, repo.db, "Expected the db to not be modified by NewBookRepository")
}
