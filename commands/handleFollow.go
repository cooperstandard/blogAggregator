package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/cooperstandard/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func HandleFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("follow requires 1 argument")
	}

	feed, err := s.DB.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	if false {
		feed, err = s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			Name:      cmd.Args[0],
			Url:       cmd.Args[1],
		})
		if err != nil {
			return err
		}
	}

	feedFollow, err := addFeedFollow(s, user.ID, feed.ID)
	if err != nil {
		return err
	}
	fmt.Printf("created feed follow record %s for user %s\n", feedFollow.FeedName, feedFollow.UserName)

	return nil
}

func addFeedFollow(s *State, userID uuid.UUID, feedID uuid.UUID) (database.CreateFeedFollowRow, error) {
	ff, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	})
	return ff, err
}
