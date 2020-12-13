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
