package commands

import (
	"fmt"
	"context"
	"time"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/rss"
)

func scrapeFeeds(s *State) error {

	ctx := context.Background()

	feed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("get next feed: %w", err)
	}
	err = s.Db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("mark fetched: %w", err)
	}

	data, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("fetch feed: %w", err)
	}
	fmt.Println("ITEMS:", len(data.Channel.Item))
	
	for _, item := range data.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}

func HandlerAgg(s *State, cmd Command) error {

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		fmt.Printf("failed to convert time with ParseDuration on HandlerAgg function: %v\n", err)
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil
}
