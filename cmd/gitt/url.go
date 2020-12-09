package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rozifus/gitt/pkg/general"
)

type UrlCmd struct {
	Url string `arg type:"url"`
}

func (urlCmd *UrlCmd) Run(ctx *general.Context) error {
	pUrl, err := url.Parse(urlCmd.Url)
	if err != nil {
		return fmt.Errorf("Invalid url: '%v'", pUrl)
	}

	switch {
	case strings.Contains(pUrl.Host, "github.com"):
		return GithubUrl(ctx, pUrl)
	case strings.Contains(pUrl.Host, "gitlab.com"):
		return GitlabUrl(ctx, pUrl)
	default:
		return fmt.Errorf("Unknown repository provider: '%s'", pUrl.Host)
	}

	return nil
}

func GithubUrl(ctx *general.Context, url *url.URL) error {
	path := strings.TrimLeft(url.Path, "/")
	spath := strings.Split(path, "/")

	switch len(spath) {
	case 1:
		s := GithubUserCmd{Username: []string{path}}
		return s.Run(ctx)
	case 2:
		s := GithubRepoCmd{UserAndRepositoryPairs: []string{path}}
		return s.Run(ctx)
	default:
		return fmt.Errorf("Couldn't figure out user or repo for github path: '%s'", path)
	}

	fmt.Println(url.Path)
	return nil
}

func GitlabUrl(ctx *general.Context, url *url.URL) error {
	fmt.Println(url.Path)
	return nil
}
