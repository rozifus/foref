package gittery

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v33/github"
	"github.com/rozifus/gitt/pkg/general"
)

// GithubUserRepositories //
func GithubUserRepositories(ctx *general.Context, usernames ...string) error {
	client := github.NewClient(nil)
	options := &github.RepositoryListOptions{Type: "owner"}
	for _, username := range usernames {
		tag := "github:" + username + "/*"
		fmt.Println(tag)
		repos, _, err := client.Repositories.List(context.Background(), username, options)
		if err != nil {
			fmt.Printf("%s : %v\n", tag, err)
		}

		downloadGithubRepositories(ctx, repos...)
	}

	return nil
}

func splitUserRepositoyPair(urp string) (username, reponame string, err error) {
	if !strings.Contains(urp, "/") {
		err = fmt.Errorf("user/repository pair must contain / but got '%s'", urp)
		return
	}

	splitUrp := strings.Split(urp, "/")

	username = splitUrp[0]
	reponame = strings.Join(splitUrp[1:], "/")
	return
}

func GithubRepositories(ctx *general.Context, userRepositoryPairs ...string) error {
	client := github.NewClient(nil)

	for _, urp := range userRepositoryPairs {
		tag := "github:" + urp
		fmt.Println(tag)

		username, reponame, err := splitUserRepositoyPair(urp)
		if err != nil {
			fmt.Printf("%s : %v\n", tag, err)
			return err
		}

		repo, _, err := client.Repositories.Get(context.Background(), username, reponame)
		if err != nil {
			fmt.Printf("%s : %v\n", tag, err)
			return err
		}

		downloadGithubRepositories(ctx, repo)
	}

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
