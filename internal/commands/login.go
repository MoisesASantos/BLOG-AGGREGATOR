package commands

import (
	"errors"
	"fmt"
	"context"
	"os"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("usage: login <username>")
	}

	ctx := context.Background()
	_, err := s.Db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		fmt.Println("The user Doesn't Exist")
		os.Exit(1)
	}

	err = s.Data.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %s logged with success\n", cmd.Args[0])
	return nil
}
