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
SET name = $1, description = $2 , is_active = $3
WHERE id = $4
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1;

-- name: ListRolesWithoutAdmin :many
SELECT * FROM roles WHERE UPPER(name) != 'ADMIN' ORDER BY id;

-- name: ListRolesWithoutAdminAndStaff :many
SELECT * FROM roles WHERE UPPER(name) NOT IN ('ADMIN', 'STAFF') ORDER BY id;


-- name: GetUserRoles :many
SELECT r.name
FROM roles r
INNER JOIN user_roles ur ON ur.role_id = r.id
WHERE ur.user_id = $1;