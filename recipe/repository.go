package recipe

import (
	"github.com/jmoiron/sqlx"
)

const (
	insertQuery = "INSERT INTO recipes VALUES (?, ?, ?, ?, ?)"
	getQuery    = "SELECT * FROM recipes WHERE id=?"
)

// sqlStore implements a recipes database using sql
type sqlStore struct {
	db *sqlx.DB
}

func (r *sqlStore) create(
	title string,
	preparationTime string,
	serves string,
	ingredients string,
	cost int,
) (*recipe, error) {
	tx := r.db.MustBegin()

	result := tx.MustExec(insertQuery, title, preparationTime, serves, ingredients, cost)

	err := tx.Commit()
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	var created recipe
	row := r.db.QueryRowx(getQuery, id)
	err = row.StructScan(&created)
	if err != nil {
		return nil, err
	}

	return &created, nil
}
