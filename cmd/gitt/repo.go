package main

import (
	"fmt"
	"github.com/rozifus/gitt/pkg/general"
)

type RepoCmd struct {
	//Read ReadCmd `cmd`
	//Test  TestCmd `cmd`
}

//type ReadCmd struct {}
//type TestCmd struct {}

func (cmd *RepoCmd) Run(ctx *general.Context) error {
	fmt.Println("Repo")
	return nil
}

/*func (cmd *ReadCmd) Run(cliCtx *CliContext) error {
	cardCtx, err := scard.EstablishContext()
	if err != nil {
		return err
	}
	defer cardCtx.Release()

	card, err := InitializeCard(cardCtx)
	if err != nil {
		return err
	}
	defer card.Disconnect(scard.ResetCard)

	res, err := silo.ReadCore(card)
	if err != nil {
		return err
	}

	printStructOfBytesAsHex(*res)

	return nil
}

func (cmd *TestCmd) Run(cliCtx *CliContext) error {
	fmt.Println("Note: SiLo must be reseated between tests, subsequent tests will fail.")

	cardCtx, err := scard.EstablishContext()
	if err != nil {
		return err
	}
	defer cardCtx.Release()

	card, err := InitializeCard(cardCtx)
	if err != nil {
		return err
	}
	defer card.Disconnect(scard.ResetCard)

	res, err := silo.TestSignature(card)
	if err != nil {
		return err
	}

	var sf string
	if (res) {
		sf = "Success"
	} else {
		sf = "Fail"
	}

	fmt.Println("Signing Test:", sf)

	return nil
}*/
