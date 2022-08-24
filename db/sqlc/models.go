// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

type Recipe struct {
	ID              int64
	Name            string
	Categories      []string
	Description     string
	PrepareMethod   string
	Ingredients     []string
	Rating          sql.NullInt32
	PreparationTime time.Time
	CreatedAt       time.Time
}

type User struct {
	ID             int64
	Email          string
	HashedPassword string
	NickName       string
	UpdatedAt      time.Time
	CreatedAt      time.Time
}
