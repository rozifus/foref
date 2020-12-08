package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/rozifus/gitt/lib/gittery"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}

	return nil
}

type GithubCmd struct {
	GithubUser GithubUserCmd `cmd name:"user" help:"Download repositories owned by a github user."`
}

type GithubUserCmd struct {
	Username string `arg`
}

func (githubUserCmd *GithubUserCmd) Run(ctx *CliContext) error {
	res, err := gittery.UserRepositories(githubUserCmd.Username)
	if err != nil {
		return err
	}

	PrettyPrint(res)

	return nil
}

type CliContext struct {
	Namespace string
}

// CLI //
var CLI struct {
	Namespace string `default:"DEFAULT" help:"Which folder namespace to use."`

	Github GithubCmd `cmd name:"github" help:"Download repositories from github."`
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run(&CliContext{Namespace: CLI.Namespace})
	ctx.FatalIfErrorf(err)
}
