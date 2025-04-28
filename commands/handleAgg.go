package commands

import (
	"context"
	"fmt"

	"github.com/cooperstandard/blogAggregator/feed"
)

func HandleAgg(s *State, cmd Command) error {
	rssFeed, err := feed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", rssFeed)

	return nil
}
