package bucket

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/rozifus/foref/pkg/source"
	"github.com/rozifus/foref/pkg/util"
)

// ListUserRepositories //
func ListUserRepositories(ctx *source.Context, identifier string) (res *bitbucket.RepositoriesRes, err error) {
	client := bitbucket.NewBasicAuth("", "")

	defer func() { recover() }()

	err = fmt.Errorf("error")
	res, err = client.Repositories.ListForAccount(&bitbucket.RepositoriesOptions{
		Owner: identifier,
	})

	return
}

// GetRepo //
func GetRepo(ctx *source.Context, identifier string) (*bitbucket.Repository, error) {
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
func DownloadRepo(ctx *source.Context, identifier string) error {
	repo, err := GetRepo(ctx, identifier)
	if err != nil {
		return err
	}

	downloadBitbucketRepositories(ctx, *repo)
	return nil
}

// UserRepositories //
func UserRepositories(ctx *source.Context, identifier string) error {
	userRepos, err := ListUserRepositories(ctx, identifier)
	if err != nil {
		return err
	}

	util.PrettyPrint(userRepos.Items)

	downloadBitbucketRepositories(ctx, userRepos.Items...)
	return nil
}

// Auto //
func Auto(ctx *source.Context, identifier string) error {
	var err error = nil
	s := strings.Split(identifier, "/")
	for len(s) > 0 {
		err = UserRepositories(ctx, strings.Join(s, "/"))
		if err == nil {
			return nil
		}

		err = DownloadRepo(ctx, strings.Join(s, "/"))
		if err == nil {
			return nil
		}

		s = s[:1]
	}

	fmt.Printf("Could not identify bitbucket user or repository for identifier '%s'", identifier)
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

func downloadBitbucketRepositories(ctx *source.Context, repositories ...bitbucket.Repository) {
	for _, repo := range repositories {
		util.PrettyPrint(repo)
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
