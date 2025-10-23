package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func FollowHandler(s *app.State, cmd Command, current_user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("not enough args for follow")
	}

	input_url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), input_url)
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    current_user.ID,
		FeedID:    feed.ID,
	}

	s.DB.CreateFeedFollow(context.Background(), params)
	return nil
}
