package main

import (
	"github.com/alecthomas/kong"
	"github.com/rozifus/gitt/pkg/general"
)

// CLI //
var CLI struct {
	Namespace string `kong:"default='DEFAULT',help='Which folder namespace to use.'"`

	/*Github    GithubCmd    `kong:"name='github',help='Download repositories from github.'"`
	Gitlab    GitlabCmd    `kong:"name='gitlab',help='Download repositories(projects) from gitlab.'"`
	Bitbucket BitbucketCmd `kong:"name='bitbucket',help='Download repositories(projects) from gitlab.'"`
	Url       UrlCmd       `kong:"name='url',help='Download repositories based on URL'"`
	*/
	Collect CollectCmd `kong:"cmd,default='1'"`
	//Identifier []string `kong:"arg"`
}

type CollectCmd struct {
	Identifier []string `kong:"arg"`
}

func (collectCmd CollectCmd) Run(ctx *general.Context) error {
	for _, identifier := range collectCmd.Identifier {
		Auto(ctx, identifier)
	}

	return nil
}

func Auto(ctx *general.Context, identifiers ...string) {

}

func main() {
	ctx := kong.Parse(&CLI)

	namespacePath, err := getNamespacePath(CLI.Namespace)
	ctx.FatalIfErrorf(err)

	generalContext := &general.Context{
		NamespacePath: namespacePath,
	}

	err = ctx.Run(generalContext)

	ctx.FatalIfErrorf(err)
}
