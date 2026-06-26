package commands

import (
	"fmt"
	"context"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.GetUserRow) error {


	ctx := context.Background()

	feed, err := s.Db.GetFeedByUrl(ctx, cmd.Args[0])
	if err != nil {
		fmt.Printf("failed to delete feed follow: %v\n", err)
		return err
	}

	
	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.Db.DeleteFeedFollow(ctx, params)

	if err != nil {
		fmt.Printf("failed to delete feed follow: %v\n", err)
		return err
	}
	return nil
}
