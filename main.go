package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/cooperstandard/blogAggregator/commands"
	"github.com/cooperstandard/blogAggregator/internal/config"
	"github.com/cooperstandard/blogAggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	s := commands.State{Config: &cfg}
	// fmt.Printf("Read config: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	queries := database.New(db)

	s.DB = queries

	cmds := commands.Commands{
		Handlers: make(map[string]func(s *commands.State, cmd commands.Command) error),
	}

	cmds.Register("login", commands.HandleLogin)
	cmds.Register("register", commands.HandleRegister)
	cmds.Register("reset", commands.HandleReset)

	arg := os.Args

	if len(arg) < 2 {
		log.Fatal("atleast 2 arguement must be provided")
	}
	cmd := commands.Command{
		Name: arg[1],
		Args: arg[2:],
	}

	err = cmds.Run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
