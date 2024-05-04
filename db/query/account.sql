-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: CreateAccount :one
Insert Into accounts (owner, balance, currency) values ($1, $2, $3) Returning *;

-- name: ListAccounts :many
Select * from accounts order by id LIMIT $1 Offset $2;

-- name: UpdateAccount :one
Update accounts set balance=$2 where id = $1 Returning *;

-- name: DeleteAccount :exec
Delete from accounts where id = $1;