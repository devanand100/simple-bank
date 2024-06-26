-- name: CreateAccount :one
INSERT INTO accounts (
  balance, owner , currency
) VALUES (
  $1, $2 , $3
)
RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
  set balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY owner
LIMIT $1 
OFFSET $2 ;

-- name: AddAmountToAccount :one
UPDATE accounts
  SET balance =  balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;