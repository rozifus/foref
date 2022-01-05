package gittfile

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


type CreateCmd struct {
	SurveyPath string `kong:"arg,type='path'"`
	IdentifierFile string `kong:"arg,type='path'"`
}

func (cmd *CreateCmd) Run(ctx *cmd.Context) error {
	rawIdentifiers := surveyInventory(cmd.SurveyPath)

	stdIdentifiers := make([]string, 0, len(rawIdentifiers))
	for _,ri := range rawIdentifiers {
		s, r, err := runner.ExtractSource(ri)

		if err != nil {
			continue
		}

		stdIdentifiers = append(stdIdentifiers, fmt.Sprintf("%s:%s", s, r))
	}


	err := createInventory(cmd.IdentifierFile, stdIdentifiers)

	return err
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


