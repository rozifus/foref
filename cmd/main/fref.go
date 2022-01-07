package main

import (
	"github.com/alecthomas/kong"

	"github.com/rozifus/foref/cmd"
	"github.com/rozifus/foref/cmd/command"
)

// CommandLine //
type CommandLine struct {
	//Source     string   `kong:"flag,short='s',optional,enum='h,github,github.com,l,gitlab,gitlab.com,b,bitbucket,bitbucket.org',default='github'"`
	//IdentifierFile string `kong:"flag,short='f',optional,type='path',help='A yaml file containing repository identifiers'"`
	//Identifier []string `kong:"arg,optional"`

	Get command.GetCmd `cmd`
	Import command.ImportCmd `cmd`
	Export command.ExportCmd `cmd`
}

// Run //
func (commandLine *CommandLine) Run() error {
	//fmt.Println("Main")
	return nil
}

func main() {
	commandLine := &CommandLine{}

	ktx := kong.Parse(commandLine)
	ktx.Run(&cmd.Context{})

	/*err := commandLine.Run()
	ktx.FatalIfErrorf(err)*/
}
