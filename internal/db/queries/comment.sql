-- name: CreateComment :one
INSERT INTO comments(comment_id, user_id, parent_id, content, type)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetCommentById :one
SELECT *
FROM comments
WHERE comment_id = $1;

-- name: GetListOfComment :many
SELECT *
FROM comments
WHERE parent_id = $1;

-- name: UpdateComment :one
UPDATE comments
SET content = $2
WHERE comment_id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE
FROM comments
WHERE comment_id = $1;