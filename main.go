package main

import (
	"fmt"
	"os"

	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/commands"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/config"
)

func main() {
	cfg := config.Read()

	state := &commands.State{
		Data: &cfg,
	}

	cmds := &commands.Commands{
		CommandMap: make(map[string]func(*commands.State, commands.Command) error),
	}

	cmds.Register("login", commands.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("usage: blog-aggregator <command> [args]")
		os.Exit(1)
	}

	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err := cmds.Run(state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
