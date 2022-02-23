-- name: GetUserFollower :many
SELECT following_user_id
from follows
WHERE followed_user_id = (@user_id)
  AND active = true;

-- name: ToggleFollow :one
INSERT INTO follows(following_user_id, followed_user_id, active)
VALUES ($1, $2, true)
ON CONFLICT ("following_user_id", "followed_user_id") DO UPDATE
    SET active = NOT (SELECT active FROM follows WHERE following_user_id = $1 AND followed_user_id = $2)
RETURNING *;