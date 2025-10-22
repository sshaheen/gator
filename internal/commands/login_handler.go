package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/sshaheen/gator/internal/app"
)

func LoginHandler(s *app.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login handler expects a single argument, the username")
	}

	username := cmd.Args[0]

	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("user %s does not exist in db\n", username)
		os.Exit(1)
	}

	if err := s.CFG.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("User name has been set to %s\n", username)

	return nil
}
