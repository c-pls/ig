-- name: GetLikeByParentId :many
SELECT *
FROM likes
WHERE parent_id = $1
  AND active = True
ORDER BY created_at DESC;

-- name: ToggleLike :one
INSERT INTO likes(parent_id, user_id, type, active)
VALUES ($1, $2, $3, true)
ON CONFLICT ("parent_id", "user_id") DO UPDATE
    SET active = NOT (SELECT active FROM likes WHERE parent_id = $1 AND user_id = $2)
RETURNING *;