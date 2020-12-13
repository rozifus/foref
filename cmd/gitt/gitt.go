package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/rozifus/gitt/pkg/general"
)

// CommandLine //
type CommandLine struct {
	Namespace  string   `kong:"flag,short='n',default='DEFAULT',help='Which folder namespace to use.'"`
	Source     string   `kong:"flag,short='s',optional,enum='github,gitlab,bitbucket',default='github'"`
	Identifier []string `kong:"arg"`
}

// Run //
func (commandLine *CommandLine) Run() error {
	namespacePath, err := getNamespacePath(commandLine.Namespace)
	if err != nil {
		return err
	}

	for _, sourceAndIdentifier := range commandLine.Identifier {
		source, identifier, err := ExtractSource(sourceAndIdentifier)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		if source == "" {
			source = commandLine.Source
		}

		ctx := &general.Context{
			NamespacePath: namespacePath,
			Source:        source,
		}

		sourceRunner(ctx, identifier)
	}

	return nil
}

func main() {
	commandLine := &CommandLine{}

	ktx := kong.Parse(commandLine)

	err := commandLine.Run()
	ktx.FatalIfErrorf(err)
}
