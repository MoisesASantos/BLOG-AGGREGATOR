package commands

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("usage: login <username>")
	}

	err := s.Data.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %s logged with success\n", cmd.Args[0])
	return nil
}
