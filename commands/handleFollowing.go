package commands

import (
	"context"
	"fmt"

	"github.com/cooperstandard/blogAggregator/internal/database"
)

func HandleFollowing(s *State, cmd Command, user database.User) error {
	feeds, err := s.DB.GetFeedFollowsByUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for i, feed := range feeds {
		fmt.Printf("%v: %s\n", i+1, feed.Name_2)
	}

	return nil
}
