package commands

import (
	"context"
	"fmt"
)

func HandleList(s *State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}
	return nil
}
