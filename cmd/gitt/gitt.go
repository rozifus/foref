package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gittbucket"
	"github.com/rozifus/gitt/pkg/gitthub"
	"github.com/rozifus/gitt/pkg/gittlab"
)

// CommandLine //
type CommandLine struct {
	Namespace  string   `kong:"flag,short='n',default='DEFAULT',help='Which folder namespace to use.'"`
	Source     string   `kong:"flag,short='s',optional,enum='github,gitlab,bitbucket,mixed',default='mixed'"`
	Identifier []string `kong:"arg"`
}

func (commandLine *CommandLine) Run() error {
	namespacePath, err := getNamespacePath(commandLine.Namespace)
	if err != nil {
		return err
	}

	for _, rawIdentifier := range commandLine.Identifier {
		source, identifier, err := ExtractIdentifierMeta(rawIdentifier)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		if source == "" {
			if commandLine.Source == "mixed" {
				fmt.Printf("Cannot determine source for identifier '%s'\n", rawIdentifier)
				continue
			}
			source = commandLine.Source
		}

		ctx := &general.Context{
			NamespacePath: namespacePath,
			Source:        source,
		}

		Auto(ctx, identifier)
	}

	return nil
}

func Auto(ctx *general.Context, identifier string) error {
	fmt.Printf("Collecting %s:%s\n", ctx.Source, identifier)
	switch ctx.Source {
	case "github.com":
		gitthub.Auto(ctx, identifier)
		return nil
	case "bitbucket.org":
		gittbucket.Auto(ctx, identifier)
		return nil
	case "gitlab.com":
		gittlab.Auto(ctx, identifier)
		return nil
	default:
		return nil
	}
}

func main() {
	commandLine := &CommandLine{}

	ktx := kong.Parse(commandLine)

	err := commandLine.Run()
	ktx.FatalIfErrorf(err)
}
