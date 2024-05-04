-- name: InsertEntry :one
Insert into entries (account_id, amount) values ($1, $2) Returning *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetEntryViaAccId :one
select * from entries where account_id = $1;

-- name: UpdateEntry :one
Update entries set amount = $2 where id = $1 Returning *;

