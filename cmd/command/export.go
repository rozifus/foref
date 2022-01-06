package command

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/go-git/go-git/v5"
	"github.com/dariubs/uniq"

	"github.com/rozifus/gitt/cmd"
	"github.com/rozifus/gitt/cmd/runner"
)


type ExportCmd struct {
	Namespace  string     `kong:"flag,short='n',default='DEFAULT',help='Which folder namespace to use.'"`
	Path  string   		  `kong:"flag,short='p',help='Which folder path to use.'"`
	IdentifierFile string `kong:"arg,type='path'"`
}

func (cmd *ExportCmd) Run(ctx *cmd.Context) error {
	path, err := runner.GetNamespacePath(cmd.Namespace, cmd.Path)
	if err != nil {
		return err
	}

	rawIdentifiers := surveyInventory(path)

	stdIdentifiers := make([]string, 0, len(rawIdentifiers))
	for _,ri := range rawIdentifiers {
		s, r, err := runner.ExtractSource(ri)

		if err != nil {
			continue
		}

		stdIdentifiers = append(stdIdentifiers, fmt.Sprintf("%s:%s", s, r))
	}

	return createInventory(cmd.IdentifierFile, stdIdentifiers)
}

func surveyInventory(path string) []string {
	identifiers := make([]string, 0)

	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		r, err := git.PlainOpen(path)
		if err != nil {
			return nil
		}

		c, err := r.Config()
		if err != nil {
			return nil
		}

		originRemote, has_origin := c.Remotes["origin"]
		if !has_origin {
			return nil
		}

		firstUrl := originRemote.URLs[0]

		identifiers = append(identifiers, firstUrl)

		return nil
	})

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	return uniq.UniqString(identifiers)
}


func createInventory(path string, identifiers []string) error {
	gd := GittfileData{
		Identifiers: identifiers,
	}

	gdYaml, err := yaml.Marshal(&gd)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, gdYaml, 0644)
	return err
}


