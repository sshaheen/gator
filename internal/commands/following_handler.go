package commands

import (
	"context"
	"fmt"

	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func FollowingHandler(s *app.State, cmd Command, current_user database.User) error {
	feed_follows, err := s.DB.GetFeedFollowsForUser(context.Background(), current_user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feed_follows {
		fmt.Println(feed.FeedName)
	}

	return nil
}
