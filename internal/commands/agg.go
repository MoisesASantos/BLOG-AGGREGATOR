package commands

import (
	"fmt"
	"context"
	"time"
	"strings"
	"database/sql"
	"github.com/google/uuid"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/rss"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
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

	for _, item := range data.Channel.Item {

		publishedAt := sql.NullTime{}

		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
				publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		err := s.Db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})

		if err != nil {
			// IGNORAR duplicados
			if strings.Contains(err.Error(), "duplicate") {
				continue
			}
			fmt.Println("error inserting post:", err)
		}	
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
	defer ticker.Stop()

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil
}
