-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY name;

-- name: CreateCategory :one
INSERT INTO categories (
  name, slug, author, isActive
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: FetchCategories :many
SELECT * FROM categories
WHERE isDeleted = false
LIMIT $1 OFFSET $2;

-- name: DeleteCategories :exec
DELETE FROM categories
WHERE id = $1;
