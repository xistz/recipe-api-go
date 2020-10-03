package main

import (
	"database/sql"
	"errors"
)

// Store defines recipe store interface
type Store interface {
	Ping() error
	FindRecipe(id int64) (*Recipe, error)
	CreateRecipe(
		title string,
		preparationTime string,
		serves string,
		ingredients string,
		cost int,
	) (int64, error)
	DeleteRecipe(id int64) error
	UpdateRecipe(
		id int64,
		title string,
		preparationTime string,
		serves string,
		ingredients string,
		cost int,
	) error
	ListRecipes() ([]*Recipe, error)
}

// NewMySQLStore creates new recipe store using mysql
func NewMySQLStore(db *sql.DB) Store {
	return &mysqlStore{db}
}

type mysqlStore struct {
	db *sql.DB
}

const (
	insertQuery = "INSERT INTO recipes(title, making_time, serves, ingredients, cost) VALUES(?, ?, ?, ?, ?)"
	findQuery   = "SELECT * FROM recipes where id=?"
	updateQuery = "UPDATE recipes SET title=?, making_time=?, serves=?, ingredients=?, cost=? WHERE id=?"
	deleteQuery = "DELETE FROM recipes where id=?"
	listQuery   = "SELECT * FROM recipes"
)

var (
	errRecipeNotFound = errors.New("recipe not found")
)

func (s *mysqlStore) Ping() error {
	err := s.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

// FindRecipe finds a recipe in the db with the given id
func (s *mysqlStore) FindRecipe(id int64) (*Recipe, error) {
	var recipe Recipe

	err := s.db.QueryRow(findQuery, id).Scan(
		&recipe.ID,
		&recipe.Title,
		&recipe.PreparationTime,
		&recipe.Serves,
		&recipe.Ingredients,
		&recipe.Cost,
		&recipe.CreatedAt,
		&recipe.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &recipe, nil
}

// CreateRecipe creates a recipe in the db with tittle, preparation time, serves, ingredients, and cost
func (s *mysqlStore) CreateRecipe(
	title string,
	preparationTime string,
	serves string,
	ingredients string,
	cost int,
) (int64, error) {
	result, err := s.db.Exec(insertQuery, title, preparationTime, serves, ingredients, cost)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// DeleteRecipe deletes recipe at id
func (s *mysqlStore) DeleteRecipe(id int64) error {
	result, err := s.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return errRecipeNotFound
	}

	return nil
}

// UpdateRecipe updates recipe at id with new details
func (s *mysqlStore) UpdateRecipe(
	id int64,
	title string,
	preparationTime string,
	serves string,
	ingredients string,
	cost int,
) error {

	result, err := s.db.Exec(updateQuery, title, preparationTime, serves, ingredients, cost, id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		return errRecipeNotFound
	}

	return nil
}

// ListRecipes returns a slice of recipes
func (s *mysqlStore) ListRecipes() ([]*Recipe, error) {
	recipes := make([]*Recipe, 0)

	rows, err := s.db.Query(listQuery)
	if err != nil {
		return recipes, err
	}

	for rows.Next() {
		var recipe Recipe

		err := rows.Scan(
			&recipe.ID,
			&recipe.Title,
			&recipe.PreparationTime,
			&recipe.Serves,
			&recipe.Ingredients,
			&recipe.Cost,
			&recipe.CreatedAt,
			&recipe.UpdatedAt,
		)
		if err != nil {
			return recipes, err
		}

		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}
