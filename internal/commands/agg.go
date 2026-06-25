package commands

import (
	"fmt"
	"context"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {

	ctx := context.Background()
	
	data, err := rss.FetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		fmt.Println("Error trying to get a fetch")
		return err
	}
	fmt.Println(data)
	return nil
}
