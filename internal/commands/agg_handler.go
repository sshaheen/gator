package commands

import (
	"context"
	"fmt"

	"github.com/sshaheen/gator/internal/app"
)

func AggregateHandler(s *app.State, cmd Command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	fetched_data, err := fetchFeed(context.Background(), feedURL)

	if err != nil {
		return err
	}

	fmt.Println(fetched_data)
	return nil
}
