package main

import (
	"github.com/alecthomas/kong"
	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gittery"
	"github.com/rozifus/gitt/pkg/gittnamespace"
)

type GithubCmd struct {
	GithubUser GithubUserCmd `cmd name:"user" help:"Download repositories owned by a github user."`
}

type GithubUserCmd struct {
	Username string `arg`
}

func (githubUserCmd *GithubUserCmd) Run(ctx *general.Context) error {
	_, err := gittery.GithubUserRepositories(ctx, githubUserCmd.Username)
	if err != nil {
		return err
	}

	//util.PrettyPrint(res)

	return nil
}

// CLI //
var CLI struct {
	Namespace string `default:"DEFAULT" help:"Which folder namespace to use."`

	Github GithubCmd `cmd name:"github" help:"Download repositories from github."`
}

func main() {
	ctx := kong.Parse(&CLI)

	namespacePath, err := gittnamespace.GetNamespacePath(CLI.Namespace)
	ctx.FatalIfErrorf(err)

	generalContext := &general.Context{
		NamespacePath: namespacePath,
	}

	err = ctx.Run(generalContext)

	ctx.FatalIfErrorf(err)
}
