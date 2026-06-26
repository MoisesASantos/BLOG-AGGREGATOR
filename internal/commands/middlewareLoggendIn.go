package commands

import (
	"context"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
)

func MiddlewareLoggedIn(
	handler func(*State, Command, database.GetUserRow) error,
) func(*State, Command) error {

	return func(s *State, cmd Command) error {

		user, err := s.Db.GetUser(
			context.Background(),
			s.Data.Current_user_name,
		)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
