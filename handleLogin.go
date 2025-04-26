package main

import "fmt"

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("please provide username as arguement after command")
	}

	s.config.SetUser(cmd.args[0])
	fmt.Printf("set user to: %s \n", cmd.args[0])
	return nil
}
