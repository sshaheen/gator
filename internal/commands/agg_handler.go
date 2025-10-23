package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/sshaheen/gator/internal/app"
)

func AggregateHandler(s *app.State, cmd Command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("agg needs 1 arg")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()

	scrapeFeeds(s, cmd)
	for ; ; <-ticker.C {
		scrapeFeeds(s, cmd)
	}
}
