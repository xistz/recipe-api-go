package main

import (
	"database/sql"
	"errors"
)

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

// FindRecipe finds a recipe in the db with the given id
func FindRecipe(db *sql.DB, id int64) (*Recipe, error) {
	var recipe Recipe

	err := db.QueryRow(findQuery, id).Scan(
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
) (int64, error) {
	result, err := db.Exec(insertQuery, title, preparationTime, serves, ingredients, cost)
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
func DeleteRecipe(db *sql.DB, id int64) error {
	result, err := db.Exec(deleteQuery, id)
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
func UpdateRecipe(
	db *sql.DB,
	id int64,
	title string,
	preparationTime string,
	serves string,
	ingredients string,
	cost int,
) error {

	result, err := db.Exec(updateQuery, title, preparationTime, serves, ingredients, cost, id)
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
func ListRecipes(db *sql.DB) ([]*Recipe, error) {
	recipes := make([]*Recipe, 0)

	rows, err := db.Query(listQuery)
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
