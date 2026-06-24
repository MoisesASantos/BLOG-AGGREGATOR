package commands

import (
	"fmt"
	"context"
	"errors"
)

func ListUsers(s *State, cmd Command) error {

	ctx := context.Background()
	users, err := s.Db.GetUsers(ctx)
	if err != nil {
		return errors.New("Something gone wrong")
	}


	for _, user := range users {
		if user.Name == s.Data.Current_user_name {
			fmt.Printf("* %s (current)\n", user.Name)

		} else
		{
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
