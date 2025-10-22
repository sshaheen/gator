package commands

import (
	"context"
	"fmt"

	"github.com/sshaheen/gator/internal/app"
)

func GetUsersHandler(s *app.State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("encountered error when fetching users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.CFG.CurrentUserName {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
			fmt.Println(user.Name)
		}
	}
	return nil
}
