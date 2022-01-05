package gittfile

import (

	"github.com/rozifus/gitt/cmd"
)


type InvCmd struct {
	Create CreateCmd `cmd`
	Load  LoadCmd `cmd`
}

func (cmd *InvCmd) Run(ctx *cmd.Context) error {
	return nil
}
