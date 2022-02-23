-- name: CreateUser :one
INSERT INTO users(user_id, username, salted_password, first_name, last_name, bio, avatar_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE user_id = $1;

-- name: GetUserByUserName :one
SELECT *
FROM users
WHERE username = $1;

-- name: UpdateFirstName :one
UPDATE users
SET first_name = $2
WHERE user_id = $1
RETURNING *;

-- name: UpdateLastName :one
UPDATE users
SET last_name = $2
WHERE user_id = $1
RETURNING *;

-- name: UpdateBio :one
UPDATE users
SET bio = $2
WHERE user_id = $1
RETURNING *;

-- name: UpdatePassword :one
UPDATE users
SET salted_password = $2
WHERE user_id = $1
RETURNING *;

-- name: UpdateAvatar :one
UPDATE users
SET avatar_url = $2
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE user_id = $1;

