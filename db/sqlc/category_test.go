package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vinicius-gregorio/receita_na_mao/util"
)

func createRandomCategory(t *testing.T) {
	categoryName := util.RandomName()

	category, err := testQueries.CreateCategory(context.Background(), categoryName)
	require.NoError(t, err)

	require.NotEmpty(t, category)
	require.Equal(t, categoryName, category.Name)
	require.NotZero(t, category.CreatedAt)
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}
