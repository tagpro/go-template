-- Example SQLC queries
-- This file demonstrates how to write SQLC queries
-- Replace with your actual queries after creating your schema

-- Example: Create a user
-- name: CreateUser :one
-- INSERT INTO users (
--     name,
--     email
-- ) VALUES (
--     $1, $2
-- ) RETURNING *;

-- Example: Get a user by ID
-- name: GetUserByID :one
-- SELECT * FROM users
-- WHERE id = $1 LIMIT 1;

-- Example: List all users
-- name: ListUsers :many
-- SELECT * FROM users
-- ORDER BY created_at DESC;

-- Example: Update a user
-- name: UpdateUser :one
-- UPDATE users
-- SET name = $2, email = $3
-- WHERE id = $1
-- RETURNING *;

-- Example: Delete a user
-- name: DeleteUser :exec
-- DELETE FROM users
-- WHERE id = $1;
