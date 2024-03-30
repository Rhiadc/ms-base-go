package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Rhiadc/ms-base-go/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type SQLStore struct {
	db         dbClient
	statement  map[string]*sql.Stmt
	migrations *migrate.Migrate
}

type dbClient interface {
	Prepare(query string) (*sql.Stmt, error)
	Close() error
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

func New(config config.Config) (*SQLStore, error) {
	db, err := sql.Open("postgres", config.DBConnectionHost)
	if err != nil {
		return nil, fmt.Errorf("database could not be initialized :%v", err)
	}

	var migrations *migrate.Migrate
	if config.MigrationsEnabled {
		migrations, err = migrate.New(config.MigrationLocation, config.DBConnectionHost)
		if err != nil {
			return nil, fmt.Errorf("migration could not be initialized :%v", err)
		}

		if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
			return nil, fmt.Errorf("migration could not be applied :%v", err)
		}
	}

	store := &SQLStore{
		db:         db,
		statement:  map[string]*sql.Stmt{},
		migrations: migrations,
	}

	err = store.prepareStatements()
	if err != nil {
		return nil, err
	}

	return store, nil
}

const (
	lockStmtKey    = "lockStatement"
	tryLockStmtkey = "trylockstatement"
)

func (s *SQLStore) prepareStatements() error {
	var err error

	s.statement[lockStmtKey], err = s.db.Prepare("SELECT pg_advisory_xact_lock($1)")

	if err != nil {
		return err
	}

	s.statement[tryLockStmtkey], err = s.db.Prepare("SELECT pg_try_advisory_xact_lock($1)")

	if err != nil {
		return err
	}

	return nil
}
