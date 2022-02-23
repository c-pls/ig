-- name: CreatePhoto :one
INSERT INTO photos(photo_id, post_id, url)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPostPhoto :many
SELECT *
FROM photos
WHERE post_id = $1;

-- name: DeletePhoto :exec
DELETE
FROM photos
WHERE photo_id = $1;