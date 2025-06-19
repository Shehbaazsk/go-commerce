-- name: CreateUser :one
INSERT INTO users (first_name ,last_name ,email ,password_hash ,phone_number ,date_of_birth ,is_active)
VALUES ($1, $2, $3, $4, $5, $6, COALESCE($7, TRUE))
RETURNING *;


-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ListUsersPaginated :many
SELECT * FROM users
ORDER BY id DESC
LIMIT $1 OFFSET $2;


-- name: UpdateUser :one
UPDATE users 
SET
  email = COALESCE(sqlc.narg('email'), email),
  first_name = COALESCE(sqlc.narg('first_name'), first_name),
  last_name = COALESCE(sqlc.narg('last_name'), last_name),
  phone_number = COALESCE(sqlc.narg('phone_number'), phone_number),
  date_of_birth = COALESCE(sqlc.narg('date_of_birth'), date_of_birth),
  is_active = COALESCE(sqlc.narg('is_active'), is_active),
  updated_by = COALESCE(sqlc.narg('updated_by'), updated_by)
WHERE id = sqlc.arg('id')
RETURNING *;


-- name: UpdateUserAuditFields :exec
UPDATE users
SET created_by = $1,
    updated_by = $2
WHERE id = $3;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: IsEmailTakenByOtherUser :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE email = $1 AND id <> $2
);

-- name: GetUserIDByCreatedBy :one
SELECT id FROM users WHERE created_by = $1;

