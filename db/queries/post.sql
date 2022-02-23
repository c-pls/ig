-- name: CreatePost :one
INSERT INTO posts(post_id, user_id, caption, longitude, latitude)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPostById :one
SELECT *
from posts
WHERE post_id = $1;

-- name: GetAllUserPost :many
SELECT *
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetAllUserPostAsc :many
SELECT *
FROM posts
WHERE user_id = $1
  AND created_at > to_timestamp(@created_at, 'YYYY-MM-DD HH24:MI:SS:US')
ORDER BY created_at
LIMIT $2;

-- name: GetAllUserPostDesc :many
SELECT *
FROM posts
WHERE user_id = $1
  AND created_at < to_timestamp(@created_at, 'YYYY-MM-DD HH24:MI:SS:US')
ORDER BY created_at DESC
LIMIT $2;

-- name: CountUserPost :one
SELECT count(*)
from posts
WHERE user_id = $1;

-- name: UpdatePostCaption :one
UPDATE posts
SET caption = $2
WHERE post_id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE
FROM posts
WHERE post_id = $1;

