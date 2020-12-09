package gittery

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v33/github"
	"github.com/rozifus/gitt/pkg/general"
)

// GithubUserRepositories //
func GithubUserRepositories(ctx *general.Context, username string) error {
	client := github.NewClient(nil)
	options := &github.RepositoryListOptions{Type: "owner"}
	repos, _, err := client.Repositories.List(context.Background(), username, options)
	if err != nil {
		return err
	}

	downloadGithubRepositories(ctx, repos...)

	return nil
}

func downloadGithubRepositories(ctx *general.Context, repos ...*github.Repository) {
	for _, repo := range repos {
		fmt.Println("github:" + *repo.FullName)
		_, err := git.PlainClone(path.Join(ctx.NamespacePath, "github.com", *repo.FullName), false, &git.CloneOptions{
			URL:      *repo.CloneURL,
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Println(*repo.FullName, err)
		}
		fmt.Println("")
	}
}
