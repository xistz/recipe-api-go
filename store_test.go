package main

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var columns = []string{
	"id",
	"title",
	"making_time",
	"serves",
	"ingredients",
	"cost",
	"created_at",
	"updated_at",
}

func initDBMock() (*sql.DB, sqlmock.Sqlmock, *sqlmock.Rows) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	rows := mock.NewRows(columns)

	return db, mock, rows
}

func TestFindRecipe(t *testing.T) {
	db, mock, rows := initDBMock()
	defer db.Close()

	expected := Recipe{
		1,
		"Chicken Curry",
		"45 min",
		"4 people",
		"onion, chicken, seasoning",
		1000,
		time.Now(),
		time.Now(),
	}

	rows = rows.AddRow(
		expected.ID,
		expected.Title,
		expected.PreparationTime,
		expected.Serves,
		expected.Ingredients,
		expected.Cost,
		expected.CreatedAt,
		expected.UpdatedAt,
	)

	mock.ExpectQuery("^SELECT \\* FROM recipes where id=\\?").WithArgs(1).WillReturnRows(rows)
	mock.ExpectQuery("^SELECT \\* FROM recipes where id=\\?").WithArgs(1000).WillReturnRows(rows)

	t.Run("returns recipe with valid ID", func(t *testing.T) {
		got, err := FindRecipe(db, expected.ID)
		assert.NoError(t, err)

		assert.Equal(t, expected.ID, got.ID)
		assert.Equal(t, expected.Title, got.Title)
		assert.Equal(t, expected.PreparationTime, got.PreparationTime)
		assert.Equal(t, expected.Serves, got.Serves)
		assert.Equal(t, expected.Ingredients, got.Ingredients)
		assert.Equal(t, expected.Cost, got.Cost)
		assert.Equal(t, expected.CreatedAt, got.CreatedAt)
		assert.Equal(t, expected.UpdatedAt, expected.UpdatedAt)
	})

	t.Run("return nil with invalid ID", func(t *testing.T) {
		got, err := FindRecipe(db, 1000)
		assert.NoError(t, err)

		assert.Nil(t, got)
	})
}

func TestCreateRecipe(t *testing.T) {
	db, mock, rows := initDBMock()
	defer db.Close()

	t.Run("returns created recipe", func(t *testing.T) {
		expected := Recipe{
			1,
			"Chicken Curry",
			"45 min",
			"4 people",
			"onion, chicken, seasoning",
			1000,
			time.Now(),
			time.Now(),
		}

		rows = rows.AddRow(
			expected.ID,
			expected.Title,
			expected.PreparationTime,
			expected.Serves,
			expected.Ingredients,
			expected.Cost,
			expected.CreatedAt,
			expected.UpdatedAt,
		)

		mock.ExpectExec("INSERT INTO recipes").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery("SELECT \\* FROM recipes where id=\\?").WithArgs(1).WillReturnRows(rows)

		got, err := CreateRecipe(
			db,
			expected.Title,
			expected.PreparationTime,
			expected.Serves,
			expected.Ingredients,
			expected.Cost,
		)

		assert.NoError(t, err)
		assert.Equal(t, expected.ID, got.ID)
		assert.Equal(t, expected.Title, got.Title)
		assert.Equal(t, expected.PreparationTime, got.PreparationTime)
		assert.Equal(t, expected.Serves, got.Serves)
		assert.Equal(t, expected.Ingredients, got.Ingredients)
		assert.Equal(t, expected.Cost, got.Cost)
		assert.NotEmpty(t, got.CreatedAt)
		assert.NotEmpty(t, got.UpdatedAt)
	})

}
