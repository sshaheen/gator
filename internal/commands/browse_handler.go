package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func BrowseHandler(s *app.State, cmd Command, current_user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		limit, _ = strconv.Atoi(cmd.Args[0])
	}

	params := database.GetPostsForUserParams{
		UserID: current_user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error grabbing posts for user %s", current_user.Name)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
	}

	return nil
}
