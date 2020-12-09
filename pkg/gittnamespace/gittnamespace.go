package gittnamespace

import (
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strings"
)

func getHome() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func get() (string, error) {
	return GetNamespacePath("default")
}

func GetNamespacePath(namespace string) (string, error) {
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
