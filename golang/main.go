package main

import (
	"log"

	"github.com/rajasoun/ollama-poc/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
