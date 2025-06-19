-- name: CreateSellerProfile :one
INSERT INTO seller_profiles (user_id, store_name, gst_number, average_rating)
VALUES ($1, $2, $3, COALESCE($4, 0.00))
RETURNING *;

-- name: UpdateSellerProfile :one
UPDATE seller_profiles
SET 
store_name = COALESCE($1,store_name), 
gst_number = COALESCE($2,gst_number),
average_rating = COALESCE($3, average_rating)
WHERE user_id = $4
RETURNING *;

-- name: GetSellerProfileByUserID :one
SELECT * FROM seller_profiles WHERE user_id = $1;

-- name: DeleteSellerProfile :exec
DELETE FROM seller_profiles WHERE user_id = $1;

-- name: ListSellersPaginated :many
SELECT 
  u.id AS user_id,
  u.first_name,
  u.last_name,
  u.email,
  u.phone_number,
  u.date_of_birth,
  u.is_active,
  sp.store_name,
  sp.gst_number,
  sp.average_rating
FROM users u
JOIN seller_profiles sp ON u.id = sp.user_id
ORDER BY u.id DESC
LIMIT $1 OFFSET $2;


