-- name: CreateGame :one
INSERT INTO games (answer, expires_at)
VALUES ($1, $2)
RETURNING *;

-- name: GetGame :one
SELECT * FROM games
WHERE id = $1;

-- name: UpdateGame :one
UPDATE games
SET guesses = $1, solved = $2
WHERE id = $3
RETURNING *;
