-- name: CreateCustomerProfile :one
INSERT INTO customer_profiles (user_id, contact_preference)
VALUES ($1, COALESCE($2, '{}'::JSONB))
RETURNING *;

-- name: UpdateCustomerProfile :one
UPDATE customer_profiles
SET
  contact_preference = COALESCE($2, contact_preference)
WHERE user_id = $1
RETURNING *;

-- name: GetCustomerByUserID :one
SELECT 
  u.id AS user_id,
  u.first_name,
  u.last_name,
  u.email,
  u.phone_number,
  u.date_of_birth,
  u.is_active,
  u.created_at AS user_created_at,
  c.contact_preference
FROM users u
JOIN customer_profiles c ON u.id = c.user_id
WHERE u.id = $1;

-- name: DeleteCustomerProfile :exec
DELETE FROM customer_profiles WHERE user_id = $1;


-- name: ListCustomersPaginated :many
SELECT 
  u.id AS user_id,
  u.first_name,
  u.last_name,
  u.email,
  u.phone_number,
  u.date_of_birth,
  u.is_active,
  u.created_at AS user_created_at,
  c.contact_preference
FROM users u
JOIN customer_profiles c ON u.id = c.user_id
ORDER BY u.id DESC
LIMIT $1 OFFSET $2;



-- name: GetCustomerByUserIDLimited :one
SELECT 
  u.first_name,
  u.last_name,
  u.email
FROM users u
JOIN customer_profiles c ON u.id = c.user_id
WHERE u.id = $1;
