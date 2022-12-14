-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  hashed_password,
  email,
  activated,
  company_name
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, created_at;


-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;



-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  first_name = COALESCE(sqlc.narg(first_name), first_name),
  last_name = COALESCE(sqlc.narg(last_name), last_name),
  activated = COALESCE(sqlc.narg(activated), activated),
  company_name =  COALESCE(sqlc.narg(company_name), company_name),
  is_blocked =  COALESCE(sqlc.narg(is_blocked), is_blocked)
WHERE
  id = sqlc.arg(id)
RETURNING *; 