package commands

import (
	"context"
	"fmt"

	"github.com/sshaheen/gator/internal/app"
)

func FeedsHandler(s *app.State, cmd Command) error {
	res, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed_row := range res {
		fmt.Printf("Feed Name: %s, URL: %s, User Name: %s\n", feed_row.FeedName, feed_row.FeedUrl, feed_row.UserName)
	}

	return nil
}
