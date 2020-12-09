package main

import (
	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gittery"
)

type GitlabCmd struct {
	GitlabUser GitlabUserCmd `cmd name:"user" help:"Download repositories owned by a gitlab user."`
}

type GitlabUserCmd struct {
	Username string `arg`
}

func (gitlabUserCmd *GitlabUserCmd) Run(ctx *general.Context) error {
	err := gittery.GitlabUserProjects(ctx, gitlabUserCmd.Username)
	if err != nil {
		return err
	}

	//util.PrettyPrint(res)

	return nil
}
