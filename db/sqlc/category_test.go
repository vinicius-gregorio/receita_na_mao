package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vinicius-gregorio/receita_na_mao/util"
)

func createRandomCategory(t *testing.T) Category {
	categoryName := util.RandomName()

	category, err := testQueries.CreateCategory(context.Background(), categoryName)
	require.NoError(t, err)

	require.NotEmpty(t, category)
	require.Equal(t, categoryName, category.Name)
	require.NotZero(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestListCategory(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCategory(t)
	}

	categories, err := testQueries.ListCategories(context.Background())
	require.NoError(t, err)

	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}

}

func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)
	require.NotEmpty(t, category)
	err := testQueries.DeleteCategory(context.Background(), category.ID)
	require.NoError(t, err)

	deletedCategory, err := testQueries.GetCategorybyId(context.Background(), category.ID)
	require.Error(t, err)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedCategory)

}

func TestUpdateCategory(t *testing.T) {
	updatedName := "UpdatedName"
	category := createRandomCategory(t)
	require.NotEmpty(t, category)

	arg := UpdateCategoryParams{
		ID:   category.ID,
		Name: updatedName,
	}
	updatedCategory, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, updatedCategory)
	require.Equal(t, category.ID, updatedCategory.ID)
	require.Equal(t, updatedCategory.Name, updatedName)
}
