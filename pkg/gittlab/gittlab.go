package gittlab

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/rozifus/gitt/pkg/general"
	"github.com/xanzy/go-gitlab"
)

// UserProjects //
func UserProjects(ctx *general.Context, username string) error {
	client, err := gitlab.NewClient("")
	if err != nil {
		return err
	}

	options := &gitlab.ListProjectsOptions{
		Visibility: gitlab.Visibility(gitlab.PublicVisibility),
		Owned:      gitlab.Bool(true),
	}
	projects, _, err := client.Projects.ListUserProjects(username, options)
	if err != nil {
		return err
	}

	DownloadRepositories(ctx, projects...)

	return err
}

// DownloadRepositories //
func DownloadRepositories(ctx *general.Context, projects ...*gitlab.Project) {
	for _, project := range projects {
		fmt.Println("gitlab:" + project.PathWithNamespace)
		_, err := git.PlainClone(path.Join(ctx.NamespacePath, "gitlab.com", project.PathWithNamespace), false, &git.CloneOptions{
			URL:      project.HTTPURLToRepo,
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Println(project.PathWithNamespace, err)
		}
		fmt.Println("")
	}
}

// GetProject //
func GetProject(ctx *general.Context, projectIdentifier string) (*gitlab.Project, error) {
	client, err := gitlab.NewClient("")
	if err != nil {
		return nil, err
	}

	project, _, err := client.Projects.GetProject(projectIdentifier, nil)
	if err != nil {
		return nil, err
	}

	return project, err
}

// GetGroup //
func GetGroup(ctx *general.Context, groupIdentifier string) (*gitlab.Group, error) {
	client, err := gitlab.NewClient("")
	if err != nil {
		return nil, err
	}

	group, _, err := client.Groups.GetGroup(groupIdentifier, nil)
	if err != nil {
		return nil, err
	}

	return group, err
}

// DownloadGroupRepositories //
func DownloadGroupRepositories(ctx *general.Context, groupIdentifier string) error {
	client, err := gitlab.NewClient("")
	if err != nil {
		return err
	}

	options := &gitlab.ListGroupProjectsOptions{
		IncludeSubgroups: gitlab.Bool(true),
	}
	projects, _, err := client.Groups.ListGroupProjects(groupIdentifier, options)
	if err != nil {
		return err
	}

	DownloadRepositories(ctx, projects...)

	return nil
}

// Auto //
func Auto(ctx *general.Context, identifiers ...string) error {
	for _, identifier := range identifiers {
		project, _ := GetProject(ctx, identifier)
		if project != nil {
			DownloadRepositories(ctx, project)
			return nil
		}

		group, _ := GetGroup(ctx, identifier)
		if group != nil {
			DownloadGroupRepositories(ctx, group.FullPath)
			return nil
		}

		return fmt.Errorf("Cannot match gitlab project or group in url identifier: '%s'", identifier)
	}
	return nil
}
