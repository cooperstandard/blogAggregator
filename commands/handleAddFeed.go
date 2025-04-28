package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cooperstandard/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func HandleAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return errors.New("adding a feed requires 2 arguments")
	}

	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
	}
	feed, err := s.DB.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", feed)

	return nil
}
