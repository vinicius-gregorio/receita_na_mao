-- name: CreateRecipe :one
INSERT INTO recipes (
    name,
    categories,
    description,
    prepare_method,
    ingredients,
    rating,
    preparation_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;


-- name: GetRecipeById :one
SELECT * FROM recipes
WHERE id = $1 LIMIT 1;

-- name: ListRecipes :many
SELECT * FROM recipes
ORDER BY id
LIMIT $1
OFFSET $2;
;

-- name: UpdateRecipe :one
UPDATE recipes
SET
    name = $2,
    categories = $3,
    description = $4,
    prepare_method = $5,
    ingredients = $6,
    rating = $7,
    preparation_time = $8
WHERE id = $1
RETURNING *
;

-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE id = $1;