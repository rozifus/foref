package runner

import (
	"fmt"

	"github.com/rozifus/foref/pkg/source"
	"github.com/rozifus/foref/pkg/source/bucket"
	"github.com/rozifus/foref/pkg/source/hub"
	"github.com/rozifus/foref/pkg/source/lab"
)


func CollectIdentifiers(ctx *source.Context, identifiers []string) error {
	for _, identifier := range identifiers {
		sourceId, repo, err := ExtractSource(ctx.Source, identifier)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		iCtx := &source.Context{
			NamespacePath: ctx.NamespacePath,
			Source: sourceId,
		}

		err = collect(iCtx, repo)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}

	return nil
}

func collect(ctx *source.Context, repo string) error {
	switch ctx.Source {
	case "github.com":
		return hub.Auto(ctx, repo)
	case "bitbucket.org":
		return bucket.Auto(ctx, repo)
	case "gitlab.com":
		return lab.Auto(ctx, repo)
	default:
		return fmt.Errorf("unknown source '%s'", ctx.Source)
	}
}
