package gitthub

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/github"
	"github.com/rozifus/gitt/pkg/general"
)

// UserRepositories //
func UserRepositories(ctx *general.Context, username string) error {
	client := github.NewClient(nil)
	options := &github.RepositoryListOptions{
		Type:        "owner",
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var repos []*github.Repository

	i := 0
	for {
		i++
		repoPage, resp, err := client.Repositories.List(context.Background(), username, options)
		if err != nil {
			return err
		}
		repos = append(repos, repoPage...)
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	downloadGithubRepositories(ctx, repos...)
	return nil
}

func Repository(ctx *general.Context, username, reponame string) error {
	client := github.NewClient(nil)

	repo, _, err := client.Repositories.Get(context.Background(), username, reponame)
	if err != nil {
		return err
	}

	downloadGithubRepositories(ctx, repo)
	return nil
}

func downloadGithubRepositories(ctx *general.Context, repos ...*github.Repository) {
	for _, repo := range repos {
		fmt.Printf("%s:%s\n", ctx.Source, *repo.FullName)
		_, err := git.PlainClone(path.Join(ctx.NamespacePath, ctx.Source, *repo.FullName), false, &git.CloneOptions{
			URL:      *repo.CloneURL,
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("")
	}
	return
}

func Auto(ctx *general.Context, identifier string) error {
	for {
		err := AutoSplit(ctx, identifier)
		if _, ok := err.(*github.RateLimitError); ok {
			fmt.Printf("%v\n", err)
			fmt.Println("hit rate limit for github, sleeping for 2 min")
			time.Sleep(2 * time.Minute)
			continue
		}
		if err != nil {
			return err
			//fmt.Printf("%v\n", err)
		}
		fmt.Println("")
		return nil
	}
}

func AutoSplit(ctx *general.Context, identifier string) error {
	s := strings.Split(identifier, "/")
	switch len(s) {
	case 0:
		return fmt.Errorf("invalid identifier '%s'", identifier)
	case 1:
		return UserRepositories(ctx, identifier)
	default:
		return Repository(ctx, s[0], s[1])
	}
}
