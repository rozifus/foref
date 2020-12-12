package gitthub

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/github"
	"github.com/rozifus/gitt/pkg/general"
)

// UserRepositories //
func UserRepositories(ctx *general.Context, usernames ...string) []error {
	client := github.NewClient(nil)
	options := &github.RepositoryListOptions{Type: "owner"}
	errs := []error{}
	for _, username := range usernames {
		tag := "github:" + username + "/*"
		fmt.Println(tag)
		repos, _, err := client.Repositories.List(context.Background(), username, options)
		if err != nil {
			fmt.Printf("%s : %v\n", tag, err)
		}

		downloadGithubRepositories(ctx, repos...)
	}

	return errs
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

func Repositories(ctx *general.Context, username, reponame string) []error {
	client := github.NewClient(nil)

	repo, _, err := client.Repositories.Get(context.Background(), username, reponame)
	if err != nil {
		return []error{err}
	}

	return downloadGithubRepositories(ctx, repo)
}

func downloadGithubRepositories(ctx *general.Context, repos ...*github.Repository) (errs []error) {
	for _, repo := range repos {
		_, err := git.PlainClone(path.Join(ctx.NamespacePath, ctx.Source, *repo.FullName), false, &git.CloneOptions{
			URL:      *repo.CloneURL,
			Progress: os.Stdout,
		})
		if err != nil {
			errs = append(errs, err)
		}
	}
	return
}

func Auto(ctx *general.Context, identifier string) []error {
	s := strings.Split(identifier, "/")
	switch len(s) {
	case 0:
		return []error{fmt.Errorf("invalid identifier '%s'", identifier)}
	case 1:
		return UserRepositories(ctx, identifier)
	case 2:
		return Repositories(ctx, s[0], s[1])
	default:
		return []error{fmt.Errorf("")}
	}
}
