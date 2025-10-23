-- name: GetPostsForUser :many
SELECT
    p.*
FROM feed_follows AS ff
JOIN users AS u ON u.id = ff.user_id
JOIN feeds AS f ON f.id = ff.feed_id
JOIN posts as p ON p.feed_id = f.id
WHERE ff.user_id = $1
ORDER BY COALESCE(p.published_at, p.created_at) DESC
LIMIT $2;