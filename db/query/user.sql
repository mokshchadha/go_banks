-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
)
Values ($1, $2, $3, $4) Returning *;

-- name: GetUser :one
SELECT  * FROM users
WHERE username = $1 LIMIT 1;