package commands

import "context"

func HandleReset(s *State, cmd Command) error {
	return s.DB.Reset(context.Background())
}
