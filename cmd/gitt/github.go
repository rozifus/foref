package main

import (
	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gittery"
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
