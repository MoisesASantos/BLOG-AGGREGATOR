/*package internal

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

type commands struct {
	Command_map	map[string]func(*state, command)error
}

func (c *commands) run(s *state, cmd command) error {

	function, ok := c.Command_map.[cmd.Name]
	if !ok {
		return errors.New("Command doesn't exist!")
	}
	err := function(stata, cmd)
	return err
}

func (c *commands) register(name string, f func(*state, command) error) {

	c.Command_map[name] = f
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
*/
