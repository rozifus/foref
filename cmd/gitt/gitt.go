package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/rozifus/gitt/cmd"
	"github.com/rozifus/gitt/cmd/gittfile"
)

// CommandLine //
type CommandLine struct {
	Namespace  string   `kong:"flag,short='n',default='DEFAULT',help='Which folder namespace to use.'"`
	//Source     string   `kong:"flag,short='s',optional,enum='h,github,github.com,l,gitlab,gitlab.com,b,bitbucket,bitbucket.org',default='github'"`
	//IdentifierFile string `kong:"flag,short='f',optional,type='path',help='A yaml file containing repository identifiers'"`
	//Identifier []string `kong:"arg,optional"`

	Inv gittfile.InvCmd `cmd`
}

// Run //
func (commandLine *CommandLine) Run() error {
	fmt.Println("Main")
	return nil
}

func main() {
	commandLine := &CommandLine{}

	ktx := kong.Parse(commandLine)
	ktx.Run(&cmd.Context{})

	/*err := commandLine.Run()
	ktx.FatalIfErrorf(err)*/
}
