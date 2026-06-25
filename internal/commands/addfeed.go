package commands

import (
	"fmt"
	"errors"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {

	if len(cmd.Args) != 2 {
		fmt.Println("Error on args, Usage: Command nameFedd urlFeed")
		return errors.New("Error on args, Usage: Command nameFedd urlFeed")
	}
	nameFeed := cmd.Args[0]
	urlFeed := cmd.Args[1]

	ctx := context.Background()
	currentUser, err := s.Db.GetUser(ctx, s.Data.Current_user_name)
	if err != nil {
		fmt.Println("Error to get a user")
		return err
	}

	now := time.Now().UTC()

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      nameFeed,
		Url:       urlFeed,
		UserID:	   currentUser.ID,
	}

	feed, err := s.Db.CreateFeed(ctx, params)
	if err != nil {
		fmt.Println("Error to create a feed")
		return err
	}

	fmt.Println(feed)
	return nil
}
