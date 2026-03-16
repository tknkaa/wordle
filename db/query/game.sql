-- name: CreateUser :one
INSERT INTO users DEFAULT VALUES
RETURNING *;

-- name: CreateSession :one
INSERT INTO sessions (user_id)
VALUES ($1)
RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1
AND expires_at > NOW();

-- name: CreateGame :one
INSERT INTO games (user_id, answer)
VALUES ($1, $2)
RETURNING *;

-- name: GetGame :one
SELECT * FROM games
WHERE id = $1;

-- name: GetTodayGameByUser :one
SELECT * FROM games
WHERE user_id = $1
AND created_at::DATE = CURRENT_DATE;

-- name: UpdateGame :one
UPDATE games
SET guesses = $1, solved = $2
WHERE id = $3
RETURNING *;
