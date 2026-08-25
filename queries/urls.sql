-- name: CreateURL :one
INSERT INTO urls (
    short_code,
    destination_url,
    expires_at
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;
