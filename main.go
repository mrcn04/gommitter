package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {

	loadEnv()

	ctx := context.Background()
	gc := createGithubClient(ctx)

	u := os.Getenv("GH_USERNAME")
	r := os.Getenv("REPO_NAME")

	err := checkIfRepoExists(ctx, gc, u, r)
	if err != nil {
		fmt.Printf("Repo is not found %+v\n", err)
		return
	}

	commit := "Daily automated commit"
	file := "commits.txt"
	content := fmt.Sprintf("Committed on: %+v\n", time.Now().Format("2006-1-2 15:4"))

	fmt.Println(content)

	err = createCommit(ctx, gc, u, r, file, content, commit)
	if err != nil {
		fmt.Printf("Error occurred while creating a commit %+v\n", err)
		return
	}
}
