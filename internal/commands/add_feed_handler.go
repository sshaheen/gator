package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func AddFeedHandler(s *app.State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("add feed requires 2 arguments")
	}

	current_usr, err := s.DB.GetUser(context.Background(), s.CFG.CurrentUserName)
	if err != nil {
		return err
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    current_usr.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create feed with error: %v", err)
	}

	fmt.Println(feed)

	return nil
}
