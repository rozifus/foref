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
