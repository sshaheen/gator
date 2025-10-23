-- name: GetNextFeedToFetch :one
SELECT id, last_fetched_at, url
FROM feeds
ORDER BY  last_fetched_at NULLS FIRST,
          last_fetched_at ASC,
          id ASC
LIMIT 1;
  