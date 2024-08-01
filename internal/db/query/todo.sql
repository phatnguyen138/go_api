-- name: CreateTodo :one
INSERT INTO todo (title, description, due_date)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListTodos :many
SELECT * FROM todo
LIMIT $1
OFFSET $2;

-- name: GetTodo :one
SELECT * FROM todo
WHERE id = $1;