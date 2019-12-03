package main

import (
	"github.com/caiorcferreira/swapi/cmd"
	"log"
)

func main() {
	if err := cmd.StartServer(8080); err != nil {
		log.Fatalf("application has failed to start due to: %v", err)
	}
}
