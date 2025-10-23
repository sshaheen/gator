package commands

import (
	"context"

	"github.com/sshaheen/gator/internal/app"
	"github.com/sshaheen/gator/internal/database"
)

func LoggedInMiddleware(handler func(s *app.State, cmd Command, user database.User) error) func(*app.State, Command) error {
	resfunc := func(s *app.State, cmd Command) error {
		user, err := s.DB.GetUser(context.Background(), s.CFG.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
	return resfunc
}
