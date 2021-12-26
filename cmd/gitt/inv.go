package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/goccy/go-yaml"

	"github.com/rozifus/gitt/pkg/general"
	"github.com/go-git/go-git/v5"
)


type InvCmd struct {
	Create CreateCmd `cmd`
	Load  LoadCmd `cmd`
}

type CreateCmd struct {
	SurveyPath string `kong:"arg,type='path'"`
	IdentifierFile string `kong:"arg,type='path'"`
}

type LoadCmd struct {
	//IdentifierFile string `kong:"flag,short='f',optional,type='path',help='A yaml file containing repository identifiers'"`
	IdentifierFiles []string `kong:"arg,optional,type='path'"`
}

func (cmd *InvCmd) Run(ctx *general.Context) error {
	return nil
}

func (cmd *CreateCmd) Run(ctx *general.Context) error {
	fmt.Println("inv:create")

	identifiers := surveyInventory(cmd.SurveyPath)
	fmt.Printf("%v\n", identifiers)

	err := createInventory(cmd.IdentifierFile, identifiers)
	fmt.Printf("%v\n", err)

	return nil
}

func (cmd *LoadCmd) Run(ctx *general.Context) error {
	fmt.Println("inv:load")

	identifiers := make([]string, 0)
	for _, idf := range cmd.IdentifierFiles {
		identifiers = append(identifiers, readIdentifierFile(idf)...)
	}

	collectIdentifiers(ctx, identifiers)

	return nil
}


type Inventory struct {
	Version string `yaml:"version,omitempty"`
	Source string `yaml:"source,omitempty"`
	Identifiers []string `yaml:"identifiers"`
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

	return identifiers
}


func createInventory(path string, identifiers []string) error {
	inventory := Inventory{
		Identifiers: identifiers,
	}

	data, err := yaml.Marshal(&inventory)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0)
	return err
}

func readIdentifierFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v", path)
		fmt.Printf("%v", err)
		return make([]string, 0)
	}

	var inventory Inventory
	if err = yaml.Unmarshal(data, &inventory); err != nil {
		fmt.Printf("Failed to parse yaml in: %v", path)
		fmt.Printf("%v", err)
		return make([]string, 0)
	}

	return inventory.Identifiers
}

