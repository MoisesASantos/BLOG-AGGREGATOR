package commands

import (
	"fmt"
	"errors"
	"context"
)

func HandlerFollowing(s *State, cmd Command) error {

	ctx := context.Background()
	user, err := s.Db.GetUser(ctx, s.Data.Current_user_name)
	if err != nil {
		return errors.New("Error trying to get user on HandlerFollowing function")
	}

	resultFollows, err :=  s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return errors.New("Error trying to get feeds on HandlerFollowing function")
	}

	for _, value := range resultFollows {
		fmt.Println(value.FeedName)
	} 
	return nil
}
