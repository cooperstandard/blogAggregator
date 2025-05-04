package main

import (
	"context"
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
	cmds.Register("users", commands.HandleList)
	cmds.Register("agg", commands.HandleAgg)
	cmds.Register("addfeed", middlewareAuth(commands.HandleAddFeed))
	cmds.Register("feeds", commands.HandleFeeds)
	cmds.Register("follow", middlewareAuth(commands.HandleFollow))
	cmds.Register("following", middlewareAuth(commands.HandleFollowing))
	cmds.Register("unfollow", middlewareAuth(commands.HandleUnfollow))

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

func middlewareAuth(handler func(s *commands.State, cmd commands.Command, user database.User) error) func(*commands.State, commands.Command) error {
	return func(s *commands.State, cmd commands.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
