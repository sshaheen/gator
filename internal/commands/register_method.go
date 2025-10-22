package commands

import "github.com/sshaheen/gator/internal/app"

func (c *Commands) Register(name string, f func(*app.State, Command) error) {
	c.Commands[name] = f
}
