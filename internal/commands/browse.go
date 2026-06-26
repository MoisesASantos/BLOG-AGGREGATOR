package commands

import (
	"fmt"
	"context"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
	"strconv"
)

func HandlerBrowse(s *State, cmd Command) error {

	ctx := context.Background()

	limit := int32(2)

	if len(cmd.Args) == 1 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			limit = int32(n)
		}
	}

	user, err := s.Db.GetUser(ctx, s.Data.Current_user_name)
	if err != nil {
		return err
	}

	posts, err := s.Db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, p := range posts {
		fmt.Printf("\n%s\n%s\n%s\n",
			p.Title,
			p.Url,
			p.Description,
		)
	}

	return nil
}
