-- name: CreateFlute :one
INSERT INTO flutes (
  description,
  available,
  key,
  name,
  material,
  holes,
  scale,
  pictures
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetFlute :one
SELECT * FROM flutes
WHERE id = $1 LIMIT 1;

-- name: GetFluteForUpdate :one
SELECT * FROM flutes
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListFlutes :many
SELECT * FROM flutes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteFlute :exec
DELETE FROM flutes
WHERE id = $1;
