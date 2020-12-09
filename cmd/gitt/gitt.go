package main

import (
	"github.com/alecthomas/kong"
	"github.com/rozifus/gitt/pkg/general"
)

// CLI //
var CLI struct {
	Namespace string `default:"DEFAULT" help:"Which folder namespace to use."`

	Github GithubCmd `cmd name:"github" help:"Download repositories from github."`
	Gitlab GitlabCmd `cmd name:"gitlab" help:"Download repositories(projects) from gitlab."`
}

func main() {
	ctx := kong.Parse(&CLI)

	namespacePath, err := getNamespacePath(CLI.Namespace)
	ctx.FatalIfErrorf(err)

	generalContext := &general.Context{
		NamespacePath: namespacePath,
	}

	err = ctx.Run(generalContext)

	ctx.FatalIfErrorf(err)
}
