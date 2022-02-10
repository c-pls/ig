-- name: CreateLike :one
INSERT INTO likes(like_id, parent_id, user_id, type, active)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ToggleLike :one
UPDATE likes
SET active = NOT active
WHERE like_id = $1
RETURNING *;