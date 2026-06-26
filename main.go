package main

import _ "github.com/lib/pq"
import (
	"fmt"
	"os"
	"database/sql"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/commands"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/config"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
)

func main() {
	cfg := config.Read()

	state := &commands.State{
		Data: &cfg,
	}

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		fmt.Println("Problem to connect with database")
		os.Exit(1)
	}
	state.Db = database.New(db) //create a new query

	cmds := &commands.Commands{
		CommandMap: make(map[string]func(*commands.State, commands.Command) error),
	}

	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.ListUsers)
	cmds.Register("agg", commands.HandlerAgg)
	cmds.Register("feeds", commands.HandlerFeeds)
	cmds.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("follow", commands.MiddlewareLoggedIn(commands.HandlerFollow))
	cmds.Register("following", commands.MiddlewareLoggedIn(commands.HandlerFollowing))
	cmds.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandlerUnfollow))

	if len(os.Args) < 2 {
		fmt.Println("usage: blog-aggregator <command> [args]")
		os.Exit(1)
	}
	
	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	err = cmds.Run(state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
