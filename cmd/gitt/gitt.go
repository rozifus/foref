package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/rozifus/gitt/pkg/general"
	"github.com/rozifus/gitt/pkg/gitthub"
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

	errs := make([]error, 0)

	for _, rawIdentifier := range commandLine.Identifier {
		source, identifier, err := ExtractIdentifierMeta(rawIdentifier)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if source == "" {
			if commandLine.Source == "mixed" {
				errs = append(errs, fmt.Errorf("Cannot determine source for '%s'", rawIdentifier))
				continue
			}
			source = commandLine.Source
		}

		ctx := &general.Context{
			NamespacePath: namespacePath,
			Source:        source,
		}

		autoErrs := Auto(ctx, identifier)
		if err != nil {
			errs = append(errs, autoErrs...)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errorlist: %v", errs)
	}

	return nil

}

func Auto(ctx *general.Context, identifier string) []error {
	switch ctx.Source {
	case "github.com":
		return gitthub.Auto(ctx, identifier)
	}
	return []error{}
}

func main() {
	commandLine := &CommandLine{}

	ktx := kong.Parse(commandLine)

	err := commandLine.Run()
	ktx.FatalIfErrorf(err)
}
