package commands

import (
	"context"
	"fmt"

	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func UnfollowHandler(s *app.State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("please supply url for unfollow")
	}

	feed_url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), feed_url)
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	return s.DB.DeleteFeedFollow(context.Background(), params)
}
