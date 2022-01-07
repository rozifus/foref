package command

import (
	"github.com/rozifus/foref/cmd"
	"github.com/rozifus/foref/cmd/runner"
	"github.com/rozifus/foref/pkg/source"
)


type GetCmd struct {
	Namespace  string   `kong:"flag,short='n',default='DEFAULT',help='Which folder namespace to use.'"`
	Path  string   		`kong:"flag,short='p',help='Which folder path to use.'"`
	Identifier []string `kong:"arg,optional,type='string'"`
}

func (cmd *GetCmd) Run(ctx *cmd.Context) error {
	path, err := runner.GetNamespacePath(cmd.Namespace, cmd.Path)
	if err != nil {
		return err
	}

	sCtx := &source.Context{
		NamespacePath: path,
	}

	return runner.CollectIdentifiers(sCtx, cmd.Identifier)
}