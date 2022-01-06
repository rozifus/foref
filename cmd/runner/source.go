package runner

import (
	"fmt"
	"net/url"
	"strings"


)


var hostMap = map[string]([]string){
	"github.com":    []string{"h", "github", "github.com"},
	"gitlab.com":    []string{"l", "gitlab", "gitlab.com"},
	"bitbucket.org": []string{"b", "bitbucket", "bitbucket.org"},
}


func ExtractSource(defaultSource, identifier string) (sourceId, repo string, err error) {
	url, err := url.Parse(identifier)

	if err == nil && url.Host != "" {
		sourceId, repo, err = splitUrl(url)
	} else {
		sourceId, repo, err = splitIdentifier(identifier)
	}

	if err != nil {
		return "", "", err
	}

	sourceId = strings.TrimLeft(sourceId, "/")

	if sourceId == "" {
		sourceId, err = coerceSourceId(defaultSource)
		if err != nil {
			return "", "", err
		}
	}

	return
}

func splitUrl(u *url.URL) (sourceId, remaining string, err error) {
	for host := range hostMap {
		if strings.Contains(u.Host, host) {
			return host, strings.TrimLeft(u.Path, "/"), nil
		}
	}
	return "", "", fmt.Errorf("Unknown repository source: '%s'", u.Host)
}

func splitIdentifier(identifier string) (sourceId, remaining string, err error) {
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

func coerceSourceId(sourceId string) (string, error) {
	lowerSource := strings.ToLower(sourceId)

	for target, aliases := range hostMap {
		for _, alias := range aliases {
			if lowerSource == alias {
				return target, nil
			}
		}
	}

	return "", fmt.Errorf("unknown source: '%s'", sourceId)
}