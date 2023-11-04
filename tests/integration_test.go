package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Rhiadc/ms-base-go/config"
	"github.com/Rhiadc/ms-base-go/domain/book"
	repo "github.com/Rhiadc/ms-base-go/infra/db/gorm"
	"github.com/Rhiadc/ms-base-go/infra/logger"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	_ "gorm.io/driver/postgres"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "testdb"
	dbPort     = "5432"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	// Init global level logger
	// Setup and teardown of the PostgreSQL container for testing
	appCfg := config.Config{ENV: "dev", Log: config.Log{DevMode: true, Encoding: "json"}}
	l := logger.NewLogger(appCfg)
	l.InitLogger()

	container := setupPostgreSQLContainer()
	defer teardownContainer(container)

	// Run the tests
	exitCode := m.Run()

	// Cleanup resources after tests
	os.Exit(exitCode)
}

func setupPostgreSQLContainer() testcontainers.Container {
	ctx := context.Background()

	postgresCtn, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		log.Fatal(err)
	}

	host, err := postgresCtn.Host(ctx)
	if err != nil {
		log.Fatal(err)
	}

	port, err := postgresCtn.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal(err)
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port.Port(), dbUser, dbPassword, dbName)

	// Initialize the test database connection

	db, err = gorm.Open(pg.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Run the database migration
	runDatabaseMigration(db)

	return postgresCtn
}

func teardownContainer(container testcontainers.Container) {
	ctx := context.Background()

	// Stop and remove the container
	_ = container.Terminate(ctx)
}

func runDatabaseMigration(db *gorm.DB) {
	// Auto-migrate the Book model to create the necessary table
	db.AutoMigrate(repo.Book{})
}

func TestCRUD(t *testing.T) {
	bookRepo := repo.NewBookRepository(db)
	bk := book.Book{Title: "Clean Code", Author: "Uncle Bob", Pages: "322"}

	b, err := bookRepo.Create(bk)
	assert.Nil(t, err)

	bg, err := bookRepo.Get(b)
	assert.Nil(t, err)
	assert.Equal(t, b, bg)
	assert.Equal(t, bg.ID, b)

	anotherBook := book.Book{Title: "Clean Arch", Author: "Uncle Bob", Pages: "322"}
	b, err = bookRepo.Create(anotherBook)
	assert.Nil(t, err)

	bg, err = bookRepo.Get(b)
	assert.Nil(t, err)
	assert.Equal(t, b, bg)
	assert.Equal(t, bg.ID, b)

	books, err := bookRepo.GetAll()
	assert.Nil(t, err)
	assert.EqualValues(t, len(books), 2)

	err = bookRepo.Delete(bg.ID)
	assert.Nil(t, err)

	bg, err = bookRepo.Get(bg.ID)
	assert.Nil(t, err)
	assert.Nil(t, bg)

}
