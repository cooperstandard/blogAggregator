package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cooperstandard/blogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
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
