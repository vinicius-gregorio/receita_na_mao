package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vinicius-gregorio/receita_na_mao/util"
)

func createRandomRecipe(t *testing.T) Recipe {
	arg := CreateRecipeParams{
		Name:            util.RandomName(),
		Categories:      []string{util.RandomName(), util.RandomName()},
		Description:     util.RandomName(),
		PrepareMethod:   util.RandomName(),
		Ingredients:     []string{util.RandomName(), util.RandomName()},
		Rating:          sql.NullInt32{2, true},
		PreparationTime: time.Now(),
	}

	recipe, err := testQueries.CreateRecipe(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, recipe)
	require.Equal(t, arg.Name, recipe.Name)
	return recipe
}

func TestCreateRecipe(t *testing.T) {
	createRandomRecipe(t)
}

func TestListRecipes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomRecipe(t)
	}
	arg := ListRecipesParams{
		Limit:  5,
		Offset: 5,
	}
	recipes, err := testQueries.ListRecipes(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, recipes)

	for _, recipe := range recipes {
		require.NotEmpty(t, recipe)
	}

}

func TestUpdateRecipe(t *testing.T) {
	recipe := createRandomRecipe(t)
	require.NotEmpty(t, recipe)

	arg := UpdateRecipeParams{
		Name:            "receitinha",
		Categories:      []string{"Bolos", "massas"},
		Description:     "lalala, lilili",
		PrepareMethod:   "pipipi popopo",
		Ingredients:     []string{util.RandomName(), util.RandomName()},
		Rating:          sql.NullInt32{2, true},
		PreparationTime: time.Now(),
		ID:              recipe.ID,
	}
	updatedRecipe, err := testQueries.UpdateRecipe(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, updatedRecipe)
	require.Equal(t, updatedRecipe.ID, recipe.ID)
}

func TestDeleteRecipe(t *testing.T) {
	recipe := createRandomRecipe(t)

	err := testQueries.DeleteRecipe(context.Background(), recipe.ID)
	require.NoError(t, err)

	getDeletedRecipe, err := testQueries.GetRecipeById(context.Background(), recipe.ID)
	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getDeletedRecipe)
}
