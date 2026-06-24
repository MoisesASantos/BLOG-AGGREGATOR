package commands

import (
	"errors"

	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/config"
)

type State struct {
	Data *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandMap map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	function, ok := c.CommandMap[cmd.Name]
	if !ok {
		return errors.New("command doesn't exist")
	}

	return function(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CommandMap[name] = f
}
