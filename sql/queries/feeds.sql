-- name: ListFeeds :many
SELECT * 
FROM feeds;

-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1
LIMIT 1;

-- name: GetFeedFollowsByUser :many
SELECT *
FROM feed_follows ff
INNER JOIN users u on ff.user_id = u.id 
INNER JOIN feeds f on ff.feed_id = f.id
WHERE ff.user_id = $1;

