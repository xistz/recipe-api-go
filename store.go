package main

import (
	"database/sql"
	"errors"
)

const (
	insertQuery = "INSERT INTO recipes(title, making_time, serves, ingredients, cost) VALUES(?, ?, ?, ?, ?)"
	findQuery   = "SELECT * FROM recipes where id=?"
)

// FindRecipe finds a recipe in the db with the given id
func FindRecipe(db *sql.DB, id int64) (*Recipe, error) {
	var recipe Recipe

	row := db.QueryRow(findQuery, id)
	err := row.Scan(
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
func CreateRecipe(
	db *sql.DB,
	title string,
	preparationTime string,
	serves string,
	ingredients string,
	cost int,
) (*Recipe, error) {

	result, err := db.Exec(insertQuery, title, preparationTime, serves, ingredients, cost)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	created, err := FindRecipe(db, id)

	return created, nil
}
