package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/rozifus/gitt/pkg/general"
)

// CommandLine //
type CommandLine struct {
	Namespace  string   `kong:"flag,short='n',default='DEFAULT',help='Which folder namespace to use.'"`
	Source     string   `kong:"flag,short='s',optional,enum='h,github,github.com,l,gitlab,gitlab.com,b,bitbucket,bitbucket.org',default='github'"`
	IdentifierFile string `kong:"flag,short='f',optional,type='path',help='A yaml file containing repository identifiers'"`
	Identifier []string `kong:"arg,optional"`
}

// Run //
func (commandLine *CommandLine) Run() error {
	namespacePath, err := getNamespacePath(commandLine.Namespace)
	if err != nil {
		return err
	}

	var identifierList []string

	if commandLine.IdentifierFile != "" {
		identifierList = readIdentifierFile(commandLine.IdentifierFile)
	} else {
		identifierList = commandLine.Identifier
	}

	for _, sourceAndIdentifier := range identifierList {
		source, identifier, err := ExtractSource(sourceAndIdentifier)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		source, err = coerceSource(commandLine.Source)
		if err != nil {
			fmt.Printf("%v", err)
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
