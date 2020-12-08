package gittery

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v33/github"
)

// UserRepositories //
func UserRepositories(username string) ([]*github.Repository, error) {
	client := github.NewClient(nil)
	opt := &github.RepositoryListOptions{Type: "owner"}
	repos, _, err := client.Repositories.List(context.Background(), username, opt)
	if err != nil {
		return nil, err
	}

	DownloadGithubRepositories(repos...)

	return repos, err
}

func DownloadGithubRepositories(repos ...*github.Repository) {
	for _, repo := range repos {
		_, err := git.PlainClone(path.Join("/home", "rozifus", "test", "wow", *repo.FullName), false, &git.CloneOptions{
			URL:      *repo.CloneURL,
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Println(*repo.FullName, err)
		}
	}
}
