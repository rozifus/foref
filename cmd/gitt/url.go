package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rozifus/gitt/pkg/util"
)

// UrlCmd //
type UrlCmd struct {
	Url string `kong:"arg,type='url'"`
}

func ExtractIdentifierMeta(identifier string) (source, remaining string, err error) {
	url, err := url.Parse(identifier)
	if err == nil && url.Host != "" {
		source, remaining, err = ParseUrl(url)
	} else {
		source, remaining, err = splitHost(identifier)
	}

	if err != nil {
		return "", "", err
	}

	source = strings.TrimLeft(source, "/")
	return
}

var hostMap = map[string]([]string){
	"github.com":    []string{"h", "github", "github.com"},
	"gitlab.com":    []string{"l", "gitlab", "gitlab.com"},
	"bitbucket.org": []string{"b", "bitbucket", "bitbucket.org"},
}

func ParseUrl(u *url.URL) (source, remaining string, err error) {
	util.PrettyPrint(hostMap)
	util.PrettyPrint(u.Host)
	for host := range hostMap {
		if strings.Contains(u.Host, host) {
			return host, u.Path, nil
		}
	}
	return "", "", fmt.Errorf("Unknown repository source: '%s'", u.Host)
}

func splitHost(identifier string) (source, remaining string, err error) {
	s := strings.Split(identifier, ":")
	switch len(s) {
	case 0:
		return "", "", fmt.Errorf("Invaliad identifier: '%s'", identifier)
	case 1:
		return "", s[0], nil
	case 2:
		source, err := coerceSource(s[0])
		if err != nil {
			return "", "", err
		}
		return source, s[1], nil
	default:
		return "", "", fmt.Errorf("Invalid identifier, too many colons: '%s'", identifier)
	}
}

func coerceSource(source string) (string, error) {

	lowerSource := strings.ToLower(source)
	for target, aliases := range hostMap {
		for _, alias := range aliases {
			if lowerSource == alias {
				return target, nil
			}
		}
	}

	return "", fmt.Errorf("unknown source: '%s'", source)

}

/*
// Run //
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
	case strings.Contains(pUrl.Host, "bitbucket.org"):
		return nil
		//return BitbucketUrl(ctx, pUrl)
	default:
		return fmt.Errorf("Unknown repository provider: '%s'", pUrl.Host)
	}
}
*/

/*
// GithubUrl //
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
}
*/
