package main

import (
	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gittery"
)

type GithubCmd struct {
	GithubUser GithubUserCmd `cmd xor:"githubcmd" name:"user" help:"Download repositories owned by a github user."`
	GithubRepo GithubRepoCmd `cmd xor:"githubcmd" name:"repo" help:"Download repositories from github."`
}

type GithubUserCmd struct {
	Username []string `arg`
}

func (githubUserCmd *GithubUserCmd) Run(ctx *general.Context) error {
	err := gittery.GithubUserRepositories(ctx, githubUserCmd.Username...)
	if err != nil {
		return err
	}

	//util.PrettyPrint(res)

	return nil
}

type GithubRepoCmd struct {
	UserAndRepositoryPairs []string `arg`
}

func (githubRepoCmd *GithubRepoCmd) Run(ctx *general.Context) error {
	err := gittery.GithubRepositories(ctx, githubRepoCmd.UserAndRepositoryPairs...)
	if err != nil {
		return err
	}

	//util.PrettyPrint(res)

	return nil
}
