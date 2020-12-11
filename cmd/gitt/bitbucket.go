package main

import (
	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gittbucket"
)

// BitbucketCmd //
type BitbucketCmd struct {
	Identifier []string `kong:"arg"`
}

// Run //
func (bitbucketCmd *BitbucketCmd) Run(ctx *general.Context) error {
	return gittbucket.Auto(ctx, bitbucketCmd.Identifier...)
}

// BitbucketUrl //
/*
func BitbucketUrl(ctx *general.Context, url *url.URL) error {
	path := strings.TrimLeft(url.Path, "/")

	project, _ := gittlab.BitbucketGetProject(ctx, path)
	if project != nil {
		gittlab.BitbucketDownloadRepositories(ctx, project)
		return nil
	}

	group, _ := gittlab.BitbucketGetGroup(ctx, path)
	if group != nil {
		gittlab.BitbucketDownloadGroupRepositories(ctx, group.FullPath)
		return nil
	}

	return fmt.Errorf("Cannot match gitlab project or group in url path: '%s'", path)
}
*/
