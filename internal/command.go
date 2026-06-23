package internal

import (
	"errors"
)

type state struct {
	Data	*config.Config
}

type command struct {
	Name	string
	Args	[]string
}

func GetCommands() map[string]command {
	return map[string]command{
		"login": {
			Name:        "login",
			Description: "Log in the user account",
			Callback:    commandLogin,
		},
	}
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) == 0 || len(cmd.Args) > 1 {
		return errors.New("Put only the command and username: example: login username")
	}

	err := s.Data.SetUser(cmd.Args[0])
	if err != nil {
		return errors.New("Failed To Log")
	}
	fmt.Printf("User %s log with sucess\n", cmd.Args[0])
	return nil
}
