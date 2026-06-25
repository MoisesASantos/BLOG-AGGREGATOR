package commands

import (
	"fmt"
	"context"
	"time"
	"errors"
	"github.com/google/uuid"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
)

func HandlerFollow(s *State, cmd Command) error {

	ctx := context.Background()

	feed, err := s.Db.GetFeedByUrl(ctx, cmd.Args[0])
	if err != nil {
		fmt.Printf("failed to create feed follow: %v\n", err)
		return err
	}
	
	user, err := s.Db.GetUser(ctx, s.Data.Current_user_name)
	if err != nil {
		return errors.New("Error trying to get user on HandlerFollowing function")
	}

	now := time.Now().UTC()
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	fmt.Println("Current user:", user.Name)
	fmt.Println("Current user ID:", user.ID)

	fmt.Println("Feed owner:", feed.UserID)
	fmt.Println("Feed ID:", feed.ID)
	resultRow, err := s.Db.CreateFeedFollow(ctx, params)
	if err != nil {
		fmt.Printf("failed to create feed follow: %v\n", err)
		return err
	}

	fmt.Println(resultRow.FeedName)
	fmt.Println(resultRow.UserName)
	return nil
}
