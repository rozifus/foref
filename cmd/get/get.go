package get

import (
	"github.com/rozifus/gitt/cmd"
	"github.com/rozifus/gitt/cmd/runner"
	"github.com/rozifus/gitt/pkg/general"
)


type GetCmd struct {
	Identifier []string `kong:"arg,optional,type='string'"`
}

func (cmd *GetCmd) Run(ctx *cmd.Context) error {
	generalCtx := &general.Context{
		NamespacePath: ctx.Namespace,
		//Source: "github.com",
	}

	return runner.CollectIdentifiers(generalCtx, cmd.Identifier)
}