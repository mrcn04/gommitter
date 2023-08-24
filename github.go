package main

import (
	"context"
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
