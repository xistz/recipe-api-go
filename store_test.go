package main

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

const (
	mockFindQuery   = "^SELECT (.+) FROM recipes where"
	mockInsertQuery = "INSERT INTO recipes"
	mockDeleteQuery = "DELETE FROM recipes where"
	mockUpdateQuery = "UPDATE recipes"
	mockListQuery   = "^SELECT (.+) FROM recipes$"
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

func initDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	return db, mock
}

func TestFindRecipe(t *testing.T) {
	db, mock := initDBMock()
	defer db.Close()

	store := NewMySQLStore(db)

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

	rows := sqlmock.NewRows(columns).AddRow(
		expected.ID,
		expected.Title,
		expected.PreparationTime,
		expected.Serves,
		expected.Ingredients,
		expected.Cost,
		expected.CreatedAt,
		expected.UpdatedAt,
	)

	mock.ExpectQuery(mockFindQuery).WithArgs(1).WillReturnRows(rows)
	mock.ExpectQuery(mockFindQuery).WithArgs(1000).WillReturnRows(rows)

	t.Run("returns recipe with valid ID", func(t *testing.T) {
		got, err := store.FindRecipe(expected.ID)
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
		got, err := store.FindRecipe(1000)
		assert.NoError(t, err)

		assert.Nil(t, got)
	})
}

func TestCreateRecipe(t *testing.T) {
	db, mock := initDBMock()
	defer db.Close()

	store := NewMySQLStore(db)

	t.Run("returns id of created recipe", func(t *testing.T) {
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

		mock.ExpectExec(mockInsertQuery).WillReturnResult(sqlmock.NewResult(1, 1))

		got, err := store.CreateRecipe(
			expected.Title,
			expected.PreparationTime,
			expected.Serves,
			expected.Ingredients,
			expected.Cost,
		)

		assert.NoError(t, err)
		assert.NotEmpty(t, got)
	})
}

func TestDeleteRecipe(t *testing.T) {
	db, mock := initDBMock()
	defer db.Close()

	store := NewMySQLStore(db)

	mock.ExpectExec(mockDeleteQuery).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(mockDeleteQuery).WithArgs(1000).WillReturnResult(sqlmock.NewResult(0, 0))

	t.Run("returns nil with valid id", func(t *testing.T) {
		err := store.DeleteRecipe(1)
		assert.NoError(t, err)
	})
}

func TestUpdateRecipe(t *testing.T) {
	db, mock := initDBMock()
	defer db.Close()

	store := NewMySQLStore(db)

	updatedRecipe := Recipe{
		ID:              1,
		Title:           "Chicken Curry",
		PreparationTime: "45 min",
		Serves:          "3 people", // updated property
		Ingredients:     "onion, chicken, seasoning",
		Cost:            1000,
	}

	mock.ExpectExec(mockUpdateQuery).
		WithArgs(
			updatedRecipe.Title,
			updatedRecipe.PreparationTime,
			updatedRecipe.Serves,
			updatedRecipe.Ingredients,
			updatedRecipe.Cost,
			updatedRecipe.ID,
		).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(mockUpdateQuery).
		WithArgs(
			updatedRecipe.Title,
			updatedRecipe.PreparationTime,
			updatedRecipe.Serves,
			updatedRecipe.Ingredients,
			updatedRecipe.Cost,
			1000,
		).
		WillReturnResult(sqlmock.NewResult(0, 0))

	t.Run("returns nil with valid id", func(t *testing.T) {

		err := store.UpdateRecipe(
			updatedRecipe.ID,
			updatedRecipe.Title,
			updatedRecipe.PreparationTime,
			updatedRecipe.Serves,
			updatedRecipe.Ingredients,
			updatedRecipe.Cost,
		)
		assert.NoError(t, err)
	})

	t.Run("returns err with invalid id", func(t *testing.T) {

		err := store.UpdateRecipe(
			updatedRecipe.ID,
			updatedRecipe.Title,
			updatedRecipe.PreparationTime,
			updatedRecipe.Serves,
			updatedRecipe.Ingredients,
			updatedRecipe.Cost,
		)
		assert.Error(t, err)
	})
}

func TestListRecipes(t *testing.T) {
	db, mock := initDBMock()
	defer db.Close()

	store := NewMySQLStore(db)

	t.Run("returns empty slice if db is empty", func(t *testing.T) {
		rows := sqlmock.
			NewRows(columns)

		mock.ExpectQuery(mockListQuery).WillReturnRows(rows)

		recipes, err := store.ListRecipes()

		assert.NoError(t, err)
		assert.Empty(t, recipes)
	})

	t.Run("returns slice of recipe pointers if db is not empty", func(t *testing.T) {
		rows := sqlmock.
			NewRows(columns).
			AddRow(
				1,
				"Chicken Curry",
				"45 min",
				"4 people",
				"onion, chicken, seasoning",
				1000,
				time.Now(),
				time.Now(),
			).
			AddRow(
				2,
				"Rice Omelette",
				"30 min",
				"2 people",
				"onion, egg, seasoning, soy sauce",
				700,
				time.Now(),
				time.Now(),
			)

		mock.ExpectQuery(mockListQuery).WillReturnRows(rows)

		recipes, err := store.ListRecipes()

		assert.NoError(t, err)
		assert.NotEmpty(t, recipes)

	})

}
