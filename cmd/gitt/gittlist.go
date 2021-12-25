package main

import (
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

type IdentifierList struct {
	Identifiers []string `yaml:"identifiers"`
}

func readIdentifierFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v", path)
		fmt.Printf("%v", err)
		return make([]string, 0)
	}

	var identifierList IdentifierList
	if err = yaml.Unmarshal(data, &identifierList); err != nil {
		fmt.Printf("Failed to parse yaml in: %v", path)
		fmt.Printf("%v", err)
		return make([]string, 0)
	}

	return identifierList.Identifiers
}
