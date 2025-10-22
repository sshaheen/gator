package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func RegisterHandler(s *app.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("register handler expects a single argument, the username")
	}

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	_, err := s.DB.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Printf("user %s already exists\n", user.Name)
		return err
	}

	if err := s.CFG.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User %s has been registered\n", user.Name)
	fmt.Println(user)

	return nil
}
