package commands

import (
	"context"
	"fmt"
)

func HandleFeeds(s *State, cmd Command) error {
	feeds, err := s.DB.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.DB.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("(%s, %s, %s)\n", user.Name, feed.Url, feed.Name)
	}
	return nil
}
