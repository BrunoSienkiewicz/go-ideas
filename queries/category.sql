-- name: ListCategories :many
SELECT *
FROM ideas.categories;

-- name: GetCategoryById :one
SELECT *
FROM ideas.categories
WHERE category_id = $1;

-- name: CreateCategory :one
INSERT INTO ideas.categories (name)
VALUES ($1)
RETURNING category_id;

-- name: UpdateCategory :one
UPDATE ideas.categories
SET name = $2
WHERE category_id = $1;

-- name: DeleteCategory :one
DELETE FROM ideas.categories
WHERE category_id = $1;

