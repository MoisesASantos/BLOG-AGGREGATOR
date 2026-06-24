package commands

import (
	"errors"
	"fmt"
	"time"
	"context"
	"os"
	"github.com/google/uuid"
	//"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/config"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	
	if len(cmd.Args) != 1 {
		return errors.New("You have to pass a name of user you wanna register")
	}

	ctx := context.Background()
	_, err := s.Db.GetUser(ctx, cmd.Args[0])
	if err == nil {
		fmt.Println("The user with this name already exist")
		os.Exit(1)
	}


	userParam := database.CreateUserParams{}
	userParam.ID = uuid.New()
	userParam.Name = cmd.Args[0]
	userParam.UpdatedAt = time.Now().UTC()

	userResult, err := s.Db.CreateUser(ctx ,userParam)
	if err == nil {
		s.Data.SetUser(userResult.Name)
	}
	return err
}
