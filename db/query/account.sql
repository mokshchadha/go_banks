-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: CreateAccount :one
Insert Into accounts (owner, balance, currency) values ($1, $2, $3) Returning *;

-- name: ListAccounts :many
Select * from accounts order by id LIMIT $1 Offset $2;