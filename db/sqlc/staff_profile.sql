-- name: CreateStaffProfile :one
INSERT INTO staff_profiles (
    user_id, employee_id, department_id, position_id, joining_date
) VALUES (
    $1, $2, $3, $4, COALESCE($5, CURRENT_DATE)
) RETURNING *;

-- name: UpdateStaffProfile :one
UPDATE staff_profiles
SET
user_id = COALESCE($1, user_id),
employee_id = COALESCE($2, employee_id),
department_id = COALESCE($3, department_id),
position_id = COALESCE($4, position_id),
joining_date = COALESCE($5, joining_date)
WHERE user_id = $6
RETURNING *;

-- name: GetStaffProfileByUserID :one
SELECT * FROM staff_profiles WHERE user_id = $1;

-- name: DeleteStaffProfile :exec
DELETE FROM staff_profiles WHERE user_id = $1;

-- name: ListStaffsPaginated :many
SELECT 
  u.id AS user_id,
  u.first_name,
  u.last_name,
  u.email,
  u.phone_number,
  u.date_of_birth,
  u.is_active,
  s.employee_id,
  s.department_id,
  s.position_id,
  s.joining_date
FROM users u
JOIN staff_profiles s ON u.id = s.user_id
ORDER BY u.id DESC
LIMIT $1 OFFSET $2;

