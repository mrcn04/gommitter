package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v53/github"
)

func main() {

	loadEnv()

	ctx := context.Background()
	gc := createGithubClient(ctx)

	_ = &github.Repository{
		Name:    github.String("testo"),
		Private: github.Bool(true),
	}

	reps, _, err := gc.Repositories.Get(ctx, os.Getenv("GH_USERNAME"), os.Getenv("REPO_NAME"))
	if err != nil {
		fmt.Printf("Repo is not found %+v\n", err)
		return
	}
	fmt.Printf("%+v \n", *reps.Name)

	fmt.Println("hello world")
	fmt.Println(os.Getenv("GH_KEY"))
}
