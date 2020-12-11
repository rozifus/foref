package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gittlab"
)

// GitlabCmd //
type GitlabCmd struct {
	GitlabAuto GitlabAutoCmd `kong:"cmd,name='auto',help='Download repositories from gitlab.'"`
	GitlabUser GitlabUserCmd `kong:"cmd,name='user',help='Download repositories owned by a gitlab user.'"`
}

// GitlabUserCmd //
type GitlabUserCmd struct {
	Username string `kong:"arg"`
}

// GitlabAuto //
type GitlabAutoCmd struct {
	Identifier []string `kong:"arg"`
}

// Run //
func (gitlabUserCmd *GitlabUserCmd) Run(ctx *general.Context) error {
	err := gittlab.UserProjects(ctx, gitlabUserCmd.Username)
	if err != nil {
		return err
	}

	//util.PrettyPrint(res)

	return nil
}

// Run //
func (gitlabAutoCmd *GitlabAutoCmd) Run(ctx *general.Context) error {
	err := gittlab.Auto(ctx, gitlabAutoCmd.Identifier...)
	if err != nil {
		return err
	}

	//util.PrettyPrint(res)

	return nil
}

// GitlabUrl //
func GitlabUrl(ctx *general.Context, url *url.URL) error {
	path := strings.TrimLeft(url.Path, "/")

	project, _ := gittlab.GetProject(ctx, path)
	if project != nil {
		gittlab.DownloadRepositories(ctx, project)
		return nil
	}

	group, _ := gittlab.GetGroup(ctx, path)
	if group != nil {
		gittlab.DownloadGroupRepositories(ctx, group.FullPath)
		return nil
	}

	return fmt.Errorf("Cannot match gitlab project or group in url path: '%s'", path)
}
