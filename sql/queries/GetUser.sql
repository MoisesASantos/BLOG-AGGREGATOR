-- name: GetUser :one
SELECT id, name, created_at 
FROM users
WHERE name = $1;

