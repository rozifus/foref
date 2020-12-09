package gittery

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/rozifus/gitt/pkg/general"
	"github.com/xanzy/go-gitlab"
)

// GitlabUserProjects //
func GitlabUserProjects(ctx *general.Context, username string) error {
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

	GitlabDownloadRepositories(ctx, projects...)

	return err
}

func GitlabDownloadRepositories(ctx *general.Context, projects ...*gitlab.Project) {
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

func GitlabGetProject(ctx *general.Context, projectId string) (*gitlab.Project, error) {
	client, err := gitlab.NewClient("")
	if err != nil {
		return nil, err
	}

	project, _, err := client.Projects.GetProject(projectId, nil)
	if err != nil {
		return nil, err
	}

	return project, err
}

func GitlabGetGroup(ctx *general.Context, groupId string) (*gitlab.Group, error) {
	client, err := gitlab.NewClient("")
	if err != nil {
		return nil, err
	}

	project, _, err := client.Groups.GetGroup(groupId, nil)
	if err != nil {
		return nil, err
	}

	return project, err
}

func GitlabDownloadGroupRepositories(ctx *general.Context, groupId string) error {
	client, err := gitlab.NewClient("")
	if err != nil {
		return err
	}

	options := &gitlab.ListGroupProjectsOptions{
		IncludeSubgroups: gitlab.Bool(true),
	}
	projects, _, err := client.Groups.ListGroupProjects(groupId, options)
	if err != nil {
		return err
	}

	GitlabDownloadRepositories(ctx, projects...)

	return nil
}
