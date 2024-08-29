-- name: ListIdeas :many
SELECT *
FROM ideas.ideas;

-- name: GetIdeaById :one
SELECT *
FROM ideas.ideas
WHERE idea_id = $1;

-- name: CreateIdea :one
INSERT INTO ideas.ideas (name, category_id)
VALUES ($1, $2)

-- name: UpdateIdea :one
UPDATE ideas.ideas
SET name = $2, category_id = $3
WHERE idea_id = $1;

-- name: DeleteIdea :one
DELETE FROM ideas.ideas
WHERE idea_id = $1;

