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

-- name: CountURLByShortCode :one
SELECT COUNT(*)
FROM urls
WHERE short_code = $1;

-- name: GetURLByShortCode :one
SELECT destination_url, short_code
FROM urls
wHERE short_code = $1;