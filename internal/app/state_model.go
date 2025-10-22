package app

import (
	"github.com/sshaheen/gator/internal/config"
	"github.com/sshaheen/gator/internal/database"
)

type State struct {
	DB  *database.Queries
	CFG *config.Config
}
