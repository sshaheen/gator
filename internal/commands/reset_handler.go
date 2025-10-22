package commands

import (
	"context"
	"fmt"

	"github.com/sshaheen/gator/internal/app"
)

func ResetHandler(s *app.State, cmd Command) error {
	err := s.DB.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("reset failed %w", err)
	}
	fmt.Println("database reset")
	return nil
}
