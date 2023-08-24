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

	u := os.Getenv("GH_USERNAME")
	r := os.Getenv("REPO_NAME")

	_, _, err := gc.Repositories.Get(ctx, u, r)
	if err != nil {
		fmt.Printf("Repo is not found %+v\n", err)
		return
	}

	// if repo is found, add a commit
	// with command, write something to txt file
	// push
}
