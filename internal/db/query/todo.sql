-- name: CreateTodo :one
INSERT INTO todo (title, description, due_date)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListTodos :many
SELECT * FROM todo
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetTodo :one
SELECT * FROM todo
WHERE id = $1;

-- name: DeleteTodo :one
DELETE FROM todo
WHERE id = $1
RETURNING *;

-- name: UpdateTodo :one
UPDATE todo
SET (title, description, completed) = ($2, $3, $4)
WHERE id = $1
RETURNING *;