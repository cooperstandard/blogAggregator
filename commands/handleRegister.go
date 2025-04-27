package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/cooperstandard/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func HandleRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("register command requires user name to register")
	}

	args := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.DB.CreateUser(context.Background(), args)
	if err != nil {
		return err
	}
	s.Config.SetUser(user.Name)

	fmt.Printf("registered user: %v\n", user)
	return nil
}
