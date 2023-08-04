package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v53/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load the env vars")
	}

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

func createGithubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_KEY")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}
