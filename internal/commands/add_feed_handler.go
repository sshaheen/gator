package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func AddFeedHandler(s *app.State, cmd Command, current_user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("add feed requires 2 arguments")
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    current_user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create feed with error: %v", err)
	}

	ff_params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    current_user.ID,
		FeedID:    feed.ID,
	}

	s.DB.CreateFeedFollow(context.Background(), ff_params)

	fmt.Println(feed)

	return nil
}
