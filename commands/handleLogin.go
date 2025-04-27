package commands

import (
	"context"
	"fmt"
)

func HandleLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("please provide username as arguement after command")
	}

	_, err := s.DB.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	s.Config.SetUser(cmd.Args[0])
	fmt.Printf("set user to: %s \n", cmd.Args[0])
	return nil
}
