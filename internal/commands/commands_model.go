package commands

import "github.com/sshaheen/gator/internal/app"

type Commands struct {
	Commands map[string]func(*app.State, Command) error
}
