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
func GitlabUserProjects(ctx *general.Context, username string) ([]*gitlab.Project, error) {
	client, err := gitlab.NewClient("")
	if err != nil {
		return nil, err
	}

	options := &gitlab.ListProjectsOptions{
		Visibility: gitlab.Visibility(gitlab.PublicVisibility),
		Owned:      gitlab.Bool(true),
	}
	projects, _, err := client.Projects.ListUserProjects(username, options)
	if err != nil {
		return nil, err
	}

	downloadGitlabRepositories(ctx, projects...)

	return projects, err
}

func downloadGitlabRepositories(ctx *general.Context, projects ...*gitlab.Project) {
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
