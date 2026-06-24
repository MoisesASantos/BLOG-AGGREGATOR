package commands

import (
	"fmt"
	"context"
	"os"
)

func HandlerReset(s *State, cmd Command) error {

	ctx := context.Background()
	err := s.Db.DeleteUser(ctx)
	if err != nil {
		fmt.Println("Error to delete all users")
		os.Exit(1)
	}
	fmt.Println("Sucess to delete all users")
	return err
}
