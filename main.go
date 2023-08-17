package main

import (
	"context"
	"fmt"
	"os"
)

func main() {

	loadEnv()

	ctx := context.Background()
	gc := createGithubClient(ctx)

	reps, _, err := gc.Repositories.Get(ctx, os.Getenv("GH_USERNAME"), os.Getenv("REPO_NAME"))
	if err != nil {
		fmt.Printf("Repo is not found %+v\n", err)
		return
	}

	// if repo is found, add a commit
	// push

	fmt.Printf("%+v \n", *reps.Name)
}
