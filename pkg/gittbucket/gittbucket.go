package gittbucket

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/rozifus/gitt/pkg/general"
)

// GetUserRepos //
func GetUserRepos(ctx *general.Context, identifier string) (*bitbucket.RepositoriesRes, error) {
	client := bitbucket.NewBasicAuth("", "")

	res, err := client.Repositories.ListForAccount(&bitbucket.RepositoriesOptions{
		Owner: identifier,
	})
	if err != nil {
		return nil, err
	}

	//util.PrettyPrint(res)

	return res, nil
}

// GetRepo //
func GetRepo(ctx *general.Context, identifier string) (*bitbucket.Repository, error) {
	si := strings.Split(identifier, "/")

	if len(si) < 2 {
		return nil, fmt.Errorf("Need 2 identifier segments for Owner and RepoSlug in '%s' len '%v'", identifier, len(si))
	}

	options := &bitbucket.RepositoryOptions{
		Owner:    si[0],
		RepoSlug: si[1],
	}

	client := bitbucket.NewBasicAuth("", "")
	return client.Repositories.Repository.Get(options)
}

// DownloadRepo //
func DownloadRepo(ctx *general.Context, identifier string) error {
	repo, err := GetRepo(ctx, identifier)
	if err != nil {
		return err
	}

	downloadBitbucketRepositories(ctx, *repo)
	return nil
}

// DownloadUserRepositories //
func DownloadUserRepositories(ctx *general.Context, identifier string) error {
	userRepos, err := GetUserRepos(ctx, identifier)
	if err != nil {
		return err
	}
	//util.PrettyPrint(user)

	downloadBitbucketRepositories(ctx, userRepos.Items...)

	return nil
	//downloadBitbutcketRepositories(ctx, user.Items.)
}

// Auto //
func Auto(ctx *general.Context, identifiers ...string) error {
	for _, identifier := range identifiers {
		repo, err := GetRepo(ctx, identifier)
		if err == nil {
			downloadBitbucketRepositories(ctx, *repo)
			continue
		}

		userRepos, err := GetUserRepos(ctx, identifier)
		if err == nil {
			downloadBitbucketRepositories(ctx, userRepos.Items...)
			continue
		}
		println(err)

		fmt.Printf("Could not identify bitbucket user or repository for '%s'", identifier)
	}
	return nil
}

func getCloneHrefFromLinks(rawLinks map[string]interface{}) string {
	cloneLinks := rawLinks["clone"].([]interface{})
	for _, cl := range cloneLinks {
		cloneLink := cl.(map[string]interface{})
		if cloneLink["name"].(string) == "https" {
			return cloneLink["href"].(string)
		}
	}
	return ""
}

func downloadBitbucketRepositories(ctx *general.Context, repositories ...bitbucket.Repository) {
	for _, repo := range repositories {
		fmt.Println("bitbucket:" + repo.Full_name)
		repoCloneHref := getCloneHrefFromLinks(repo.Links)
		_, err := git.PlainClone(path.Join(ctx.NamespacePath, "bitbucket.org", repo.Full_name), false, &git.CloneOptions{
			URL:      repoCloneHref,
			Progress: os.Stdout,
		})

		if err != nil {
			fmt.Println(repo.Full_name, err)
		}

		fmt.Println("")
	}
}
