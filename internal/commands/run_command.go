package commands

import (
	"fmt"

	"github.com/sshaheen/gator/internal/app"
)

func (c *Commands) Run(s *app.State, cmd Command) error {
	if _, ok := c.Commands[cmd.Name]; ok {
		return c.Commands[cmd.Name](s, cmd)
	}
	return fmt.Errorf("no method for command name %s", cmd.Name)
}
