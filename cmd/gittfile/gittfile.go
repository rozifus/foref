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
	"github.com/rozifus/gitt/pkg/general"
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

func (cmd *InvCmd) Run(ctx *cmd.Context) error {
	return nil
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

func (cmd *LoadCmd) Run(ctx *cmd.Context) error {
	fmt.Println("inv:load")

	gDatas := make([]*GittfileData, 0)
	for _, gFile := range cmd.IdentifierFiles {
		gDatas = append(gDatas, readGittfile(gFile))
	}

	identifiers := make([]string, 0)
	for _, gData := range gDatas {
		identifiers = append(identifiers, gData.Identifiers...)
	}

	generalCtx := &general.Context{
		NamespacePath: ctx.NamespacePath,
		Source: "github.com",
	}

	fmt.Printf("okay\n")
	fmt.Printf("%v\n", generalCtx)

	runner.CollectIdentifiers(generalCtx, identifiers)

	return nil
}


type GittfileData struct {
	GittfileNotice string `yaml:"gittfile_notice,omitempty"`
	GittfileVersion string `yaml:"gittfile_version,omitempty"`
	Description string `yaml:"description,omitempty"`
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

func readGittfile(path string) *GittfileData {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v", path)
		fmt.Printf("%v", err)
		return nil
	}

	var gd GittfileData
	if err = yaml.Unmarshal(data, &gd); err != nil {
		fmt.Printf("Failed to parse yaml in: %v", path)
		fmt.Printf("%v", err)
		return nil
	}

	return &gd
}

