package commands

import (
	"fmt"
	"errors"
	"context"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.GetUserRow) error {

	ctx := context.Background()
	resultFollows, err :=  s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return errors.New("Error trying to get feeds on HandlerFollowing function")
	}

	for _, value := range resultFollows {
		fmt.Println(value.FeedName)
	} 
	return nil
}
