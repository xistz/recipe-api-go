package main

import "time"

// Recipe defines model for recipe
type Recipe struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	PreparationTime string    `json:"preparation_time"`
	Serves          string    `json:"serves"`
	Ingredients     string    `json:"ingredients"`
	Cost            int       `json:"cost"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}
