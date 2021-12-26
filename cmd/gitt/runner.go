package main

import (
	"fmt"

	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gittbucket"
	"github.com/rozifus/gitt/pkg/gitthub"
	"github.com/rozifus/gitt/pkg/gittlab"
)

func sourceRunner(ctx *general.Context, identifier string) error {
	fmt.Printf("Collecting %s:%s\n", ctx.Source, identifier)
	switch ctx.Source {
	case "github.com":
		return gitthub.Auto(ctx, identifier)
	case "bitbucket.org":
		return gittbucket.Auto(ctx, identifier)
	case "gitlab.com":
		return gittlab.Auto(ctx, identifier)
	default:
		return nil
	}
}

func collectIdentifiers(ctx *general.Context, identifiers []string) error {
	/*namespacePath, err := getNamespacePath(commandLine.Namespace)
	if err != nil {
		return err
	}*/

	/*var identifierList []string

	if commandLine.IdentifierFile != "" {
		identifierList = readIdentifierFile(commandLine.IdentifierFile)
	} else {
		identifierList = commandLine.Identifier
	}*/

	for _, sourceAndIdentifier := range identifiers {
		source, identifier, err := ExtractSource(sourceAndIdentifier)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		// TODO: something other than this
		if source == "" && ctx.Source == "" {
			ctx.Source = "github"
		}

		source, err = coerceSource(ctx.Source)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		ctx := &general.Context{
			NamespacePath: "",//namespacePath, TODO: fix
			Source:        source,
		}

		sourceRunner(ctx, identifier)
	}

	return nil
}
