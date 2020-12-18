package main

import (
	"fmt"
	"net/url"
	"strings"
)

func ExtractSource(identifier string) (source, remaining string, err error) {
	url, err := url.Parse(identifier)
	if err == nil && url.Host != "" {
		source, remaining, err = parseUrl(url)
	} else {
		source, remaining, err = splitHost(identifier)
	}

	if err != nil {
		return "", "", err
	}

	source = strings.TrimLeft(source, "/")
	return
}

func parseUrl(u *url.URL) (source, remaining string, err error) {
	for host := range hostMap {
		if strings.Contains(u.Host, host) {
			return host, strings.TrimLeft(u.Path, "/"), nil
		}
	}
	return "", "", fmt.Errorf("Unknown repository source: '%s'", u.Host)
}

var hostMap = map[string]([]string){
	"github.com":    []string{"h", "github", "github.com"},
	"gitlab.com":    []string{"l", "gitlab", "gitlab.com"},
	"bitbucket.org": []string{"b", "bitbucket", "bitbucket.org"},
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

func splitHost(identifier string) (source, remaining string, err error) {
	s := strings.Split(identifier, ":")
	switch len(s) {
	case 0:
		return "", "", fmt.Errorf("Invaliad identifier: '%s'", identifier)
	case 1:
		return "", s[0], nil
	case 2:
		return s[0], s[1], nil
	default:
		return "", "", fmt.Errorf("Invalid identifier, too many colons: '%s'", identifier)
	}
}
