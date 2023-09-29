package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	loadEnv()

	ctx := context.Background()
	gc := createGithubClient(ctx)

	err := checkIfRepoExists(ctx, gc)
	if err != nil {
		fmt.Printf("Repo is not found %+v\n", err)
		return
	}

	commit := "Daily automated commit"
	file := "commit.txt"
	content := fmt.Sprintf("Committed on: %+v\n", time.Now().Format("2006-1-2 15:4"))

	err = createCommit(ctx, gc, file, content, commit)
	if err != nil {
		fmt.Printf("Error occurred while creating a commit %+v\n", err)
		return
	}
}
