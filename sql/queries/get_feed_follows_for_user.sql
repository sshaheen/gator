-- name: GetFeedFollowsForUser :many
SELECT
    ff.*,
    f.name AS feed_name,
    u.name AS user_name
FROM feed_follows AS ff
LEFT JOIN users AS u ON u.id = ff.user_id
LEFT JOIN feeds AS f ON f.id = ff.feed_id
WHERE ff.user_id = $1;