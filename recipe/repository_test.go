package recipe

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func newMock() (sqlStore, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("unexpected error occurred: %s", err)
	}

	dbx := sqlx.NewDb(db, "sqlmock")
	store := sqlStore{dbx}

	return store, mock
}

func TestStoreCreate(t *testing.T) {
	store, mock := newMock()
	defer store.db.Close()

	expected := recipe{
		1,
		"Chicken Curry",
		"45 min",
		"4 people",
		"onion, chicken, seasoning",
		1000,
		time.Now(),
		time.Now(),
	}

	row := sqlmock.NewRows(
		[]string{
			"id",
			"title",
			"making_time",
			"serves",
			"ingredients",
			"cost",
			"created_at",
			"updated_at",
		},
	).AddRow(
		expected.ID,
		expected.Title,
		expected.PreparationTime,
		expected.Serves,
		expected.Ingredients,
		expected.Cost,
		expected.CreatedAt,
		expected.UpdatedAt,
	)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO recipes").WithArgs(expected.Title, expected.PreparationTime, expected.Serves, expected.Ingredients, expected.Cost).WillReturnResult(sqlmock.NewResult(expected.ID, 1))
	mock.ExpectCommit()
	mock.ExpectQuery("SELECT \\* FROM recipes WHERE id=\\?").WithArgs(expected.ID).WillReturnRows(row)

	got, err := store.create(expected.Title, expected.PreparationTime, expected.Serves, expected.Ingredients, expected.Cost)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
	assert.Equal(t, expected.ID, got.ID)
	assert.Equal(t, expected.Title, got.Title)
	assert.Equal(t, expected.PreparationTime, got.PreparationTime)
	assert.Equal(t, expected.Serves, got.Serves)
	assert.Equal(t, expected.Ingredients, got.Ingredients)
	assert.Equal(t, expected.Cost, got.Cost)
	assert.NotEmpty(t, got.CreatedAt)
	assert.NotEmpty(t, got.UpdatedAt)
}
