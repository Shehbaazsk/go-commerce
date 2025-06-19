-- name: CreateDepartment :one
INSERT INTO departments (name, manager_id)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateDepartment :one
UPDATE departments
SET 
    name = COALESCE($1, name), 
    manager_id = COALESCE($2, manager_id)
WHERE id = $3
RETURNING *;

-- name: GetDepartmentByID :one
SELECT * FROM departments WHERE id = $1;

-- name: DeleteDepartment :exec
DELETE FROM departments WHERE id = $1;

-- name: ListDepartmentsPaginated :many
SELECT * FROM departments
ORDER BY id DESC
LIMIT $1 OFFSET $2;

-- name: CreatePosition :one
INSERT INTO positions (title, description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdatePosition :one
UPDATE positions
SET 
    title = COALESCE($1, title), 
    description = COALESCE($2, description)
WHERE id = $3
RETURNING *;

-- name: GetPositionByID :one
SELECT * FROM positions WHERE id = $1;

-- name: DeletePosition :exec
DELETE FROM positions WHERE id = $1;

-- name: ListPositionsPaginated :many
SELECT * FROM positions
ORDER BY id DESC
LIMIT $1 OFFSET $2;

