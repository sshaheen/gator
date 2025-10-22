package commands

import (
	"bytes"
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	data := []byte{}
	data_reader := bytes.NewBuffer(data)
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, data_reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	rb_data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var feed *RSSFeed

	err = xml.Unmarshal(rb_data, &feed)
	if err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	return feed, nil
}
