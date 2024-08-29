-- name: ListAttributes :many
SELECT * 
FROM ideas.attributes;

-- name: GetAttributeById :one
SELECT * 
FROM ideas.attributes 
WHERE attribute_id = $1;

-- name: GetAttributesByIdeaId :many
SELECT *
FROM ideas.attributes
WHERE idea_id = $1;

-- name: GetAttributesByField :many
SELECT *
FROM ideas.attributes
WHERE $1 = $2;

-- name: CreateAttribute :one
INSERT INTO ideas.attributes (name, value, idea_id)
VALUES ($1, $2, $3) 
RETURNING attribute_id;

-- name: UpdateAttribute :one
UPDATE ideas.attributes
SET name = $2, value = $3, idea_id = $4
WHERE attribute_id = $1;

-- name: DeleteAttribute :one
DELETE FROM ideas.attributes
WHERE attribute_id = $1;

