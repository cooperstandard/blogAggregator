package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cooperstandard/blogAggregator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("command %s is not registered to handler", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}
	s := state{&cfg}

	fmt.Printf("Read config: %+v\n", cfg)

	cmds := commands{
		handlers: make(map[string]func(s *state, cmd command) error),
	}

	cmds.register("login", handleLogin)

	arg := os.Args

	if len(arg) < 2 {
		log.Fatal("atleast 2 arguement must be provided")
	}
	cmd := command{
		name: arg[1],
		args: arg[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("please provide username as arguement after command")
	}

	s.config.SetUser(cmd.args[0])
	fmt.Printf("set user to: %s \n", cmd.args[0])
	return nil
}
