-- name: CreateUserRole :one
INSERT INTO user_roles (
  user_id, role_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserRoles :many
SELECT r.name
FROM roles r
JOIN user_roles ur ON ur.role_id = r.id
WHERE ur.user_id = $1;

-- name: DeleteUserRole :exec
DELETE FROM user_roles WHERE user_id = $1 AND role_id = $2;

-- name: UpdateUserRole :one
UPDATE user_roles
SET role_id = $2
WHERE user_id = $1
RETURNING *;
