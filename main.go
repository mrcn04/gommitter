package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load the env vars")
	}

	fmt.Println("hello world")
	fmt.Println(os.Getenv("GITHUB_KEY"))
}
