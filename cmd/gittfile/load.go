package gittfile

import (
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml"

	"github.com/rozifus/gitt/cmd"
	"github.com/rozifus/gitt/cmd/runner"
	"github.com/rozifus/gitt/pkg/general"
)


type LoadCmd struct {
	//IdentifierFile string `kong:"flag,short='f',optional,type='path',help='A yaml file containing repository identifiers'"`
	IdentifierFiles []string `kong:"arg,optional,type='path'"`
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
		NamespacePath: ctx.Namespace,
		Source: "github.com",
	}

	runner.CollectIdentifiers(generalCtx, identifiers)

	return nil
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

