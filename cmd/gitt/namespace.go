package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getNamespacePath(namespace string) (string, error) {
	alphaNumericRegex := regexp.MustCompile("^[a-zA-Z0-9]+$")
	if !alphaNumericRegex.MatchString(namespace) {
		return "", fmt.Errorf("namespace should be alphanumeric but got '%s'", namespace)
	}

	envName := "GITT_NAMESPACE_" + strings.ToUpper(namespace)

	namespacePath := os.Getenv(envName)
	if namespacePath == "" {
		return "", fmt.Errorf("env variable not set for %s", envName)
	}

	return namespacePath, nil
}
