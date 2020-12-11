package main

import (
	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gitthub"
)

// GithubCmd //
type GithubCmd struct {
	GithubUser GithubUserCmd `kong:"cmd,xor='githubcmd',name='user',help='Download repositories owned by a github user.'"`
	GithubRepo GithubRepoCmd `kong:"cmd,xor='githubcmd',name='repo',help='Download repositories from github.'"`
}

// GithubUserCmd //
type GithubUserCmd struct {
	Username []string `kong:"arg"`
}

// Run //
func (githubUserCmd *GithubUserCmd) Run(ctx *general.Context) error {
	err := gitthub.GithubUserRepositories(ctx, githubUserCmd.Username...)
	if err != nil {
		return err
	}

	//util.PrettyPrint(res)

	return nil
}

// GithubRepoCmd //
type GithubRepoCmd struct {
	UserAndRepositoryPairs []string `kong:"arg"`
}

// Run //
func (githubRepoCmd *GithubRepoCmd) Run(ctx *general.Context) error {
	err := gitthub.GithubRepositories(ctx, githubRepoCmd.UserAndRepositoryPairs...)
	if err != nil {
		return err
	}

	//util.PrettyPrint(res)

	return nil
}
