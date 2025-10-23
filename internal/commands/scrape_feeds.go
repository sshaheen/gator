package commands

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func scrapeFeeds(s *app.State, cmd Command) error {
	next_feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	if err := s.DB.MarkFeedFetched(context.Background(), next_feed.ID); err != nil {
		return err
	}

	feed_data, err := fetchFeed(context.Background(), next_feed.Url)
	if err != nil {
		return err
	}

	layout := time.Layout

	for _, item := range feed_data.Channel.Item {
		publish_date, err := time.Parse(layout, item.PubDate)
		if err != nil {
			publish_date = time.Now()
		}
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publish_date,
			FeedID:      next_feed.ID,
		}

		s.DB.CreatePost(context.Background(), params)
	}

	return nil
}
