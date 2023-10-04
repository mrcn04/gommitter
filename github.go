package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func createGithubClient(ctx context.Context) *github.Client {
	t := os.Getenv("GH_KEY")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: t})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

func checkIfRepoExists(ctx context.Context, gc *github.Client) error {
	_, _, err := gc.Repositories.Get(ctx, User, Repo)
	if err != nil {
		return err
	}
	return nil
}

func createCommit(ctx context.Context, gc *github.Client, file string, content string, commit string) error {
	fmt.Println("Creating a commit...")
	// Get the default branch
	db, _, err := gc.Repositories.GetBranch(ctx, User, Repo, Branch, false)
	if err != nil {
		return err
	}

	newBlob, _, err := gc.Git.CreateBlob(ctx, User, Repo, &github.Blob{
		Content:  github.String(content),
		Encoding: github.String("utf-8"),
	})
	if err != nil {
		fmt.Println(2)
		return err
	}

	// Get the tree associated with the latest commit
	tree, _, err := gc.Git.GetTree(ctx, User, Repo, *db.Commit.SHA, true)
	if err != nil {
		fmt.Println(1)
		return err
	}

	// Add the new file entry to the existing tree
	newTreeEntries := append(tree.Entries, &github.TreeEntry{
		Path: github.String(file),
		Mode: github.String("100644"), // File mode
		Type: github.String("blob"),
		SHA:  newBlob.SHA,
	})

	// Create a new tree with the updated tree entries
	newTree, _, _ := gc.Git.CreateTree(ctx, User, Repo, *tree.SHA, newTreeEntries)

	// Create a new commit based on the new tree
	newCommit, _, err := gc.Git.CreateCommit(ctx, User, Repo, &github.Commit{
		Message: github.String(commit),
		Tree:    newTree,
		Parents: []*github.Commit{{SHA: db.Commit.SHA}},
	})
	if err != nil {
		fmt.Println(3)
		return err
	}

	// Update the reference (branch) to point to the new commit
	_, _, err = gc.Git.UpdateRef(ctx, User, Repo, &github.Reference{
		Ref: github.String(Ref),
		Object: &github.GitObject{
			Type: github.String(Type),
			SHA:  newCommit.SHA,
		},
	}, true)

	if err != nil {
		return err
	}

	fmt.Printf("Commit created: %s\n", *newCommit.Message)
	return nil
}

//
// For updating the wanted file directly using Repository API
//
// func createCommit(ctx context.Context, gc *github.Client, u string, r string, file string, content string, commit string) error {
// 	// Get the latest commit SHA of the default branch
// 	latestCommit, _, _ := gc.Repositories.GetBranch(ctx, u, r, "master", false)

// 	// Get the existing content of the file
// 	existingContent, _, _, _ := gc.Repositories.GetContents(ctx, u, r, file, &github.RepositoryContentGetOptions{
// 		Ref: *latestCommit.Commit.SHA,
// 	})

// 	// Combine existing content and new content with a newline
// 	updatedContent := *existingContent.Content + "\n" + content

// 	// Update the file content
// 	_, _, err := gc.Repositories.UpdateFile(ctx, u, r, file, &github.RepositoryContentFileOptions{
// 		Message: github.String(commit),
// 		Content: []byte(updatedContent), // Convert content to []byte
// 		SHA:     existingContent.SHA,
// 	})

// 	if err != nil {
// 		fmt.Printf("Error updating file: %v\n", err)
// 		return err
// 	}

// 	fmt.Printf("File updated successfully.\n")

// 	return nil
// }
