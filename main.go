package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	loadEnv()

	http.HandleFunc("/", hello)
	http.HandleFunc("/commit-github", handleCommit)

	p := os.Getenv("PORT")

	fmt.Printf("Starting server at port %s\n", p)
	if err := http.ListenAndServe(":"+p, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is running on %s", p)
}

func handleCommit(w http.ResponseWriter, r *http.Request) {
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
		fmt.Printf("Error occurred while creating a commit: %+v\n", err)
		return
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello There!"))
}
