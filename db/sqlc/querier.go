// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateCategory(ctx context.Context, name string) (Category, error)
	DeleteCategory(ctx context.Context, id int64) error
	GetCategorybyId(ctx context.Context, id int64) (Category, error)
	ListCategories(ctx context.Context) ([]Category, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
}

var _ Querier = (*Queries)(nil)