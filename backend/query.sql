-- Example query file
-- Replace or modify these queries according to your needs

-- Example: Get a single user by ID
-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- Example: Get a user by username
-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- Example: List users with pagination
-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- Example: Create a new user
-- name: CreateUser :one
INSERT INTO users (
    username,
    email
) VALUES (
    $1, $2
) RETURNING *;

-- Example: Update a user
-- name: UpdateUser :one
UPDATE users
SET username = $2,
    email = $3,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- Example: Delete a user
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- Add your custom queries below
-- Example:
-- name: YourCustomQuery :many
-- SELECT * FROM your_table
-- WHERE field1 = $1; 