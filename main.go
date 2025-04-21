package main

import (
	"fmt"
	"log"

	"github.com/cooperstandard/blogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}

	fmt.Printf("Read config: %+v\n", cfg)

	cfg.SetUser("lane")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
