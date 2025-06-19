-- name: CreateRole :one
INSERT INTO roles (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetRoleByID :one
SELECT * FROM roles rs WHERE id = $1;

-- name: ListRoles :many
SELECT * FROM roles ORDER BY id;

-- name: UpdateRole :one
UPDATE roles
SET 
    name        = COALESCE(sqlc.narg('name'), name),
    description = COALESCE(sqlc.narg('description'), description),
    is_active   = COALESCE(sqlc.narg('is_active'), is_active)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1;

-- name: ListRolesWithoutAdmin :many
SELECT * FROM roles WHERE UPPER(name) != 'ADMIN' ORDER BY id;

-- name: ListRolesWithoutAdminAndStaff :many
SELECT * FROM roles WHERE UPPER(name) NOT IN ('ADMIN', 'STAFF') ORDER BY id;