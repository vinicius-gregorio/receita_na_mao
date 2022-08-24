-- name: CreateRecipe :one
INSERT INTO recipes (
    name,
    categories,
    description,
    prepare_method,
    ingredients,
    rating,
    preparation_time,

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
name,
    categories = $2,
    description = $3,
    prepare_method = $4,
    ingredients = $5,
    rating = $6,
    preparation_time = $7,

WHERE id = $1
RETURNING *
;

-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE id = $1;