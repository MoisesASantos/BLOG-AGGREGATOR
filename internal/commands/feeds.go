package commands

import (
	"fmt"
	"errors"
	"context"
)

func HandlerFeeds(s *State, cmd Command) error {

	ctx := context.Background()

	resultFeed, err := s.Db.GetFeeds(ctx)
	if err != nil {
		fmt.Println("Error trying to get a feed")
		return errors.New("Error trying to get a feed")
	}

	for index, _ := range resultFeed {

		fmt.Println(resultFeed[index].Name)
		fmt.Println(resultFeed[index].Url)
		fmt.Println(resultFeed[index].Name_2)
	}
	return nil
}
