package gittnamespace

import (
	"os"
	"os/user"
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
	return getNamespaced("default")
}

func getNamespaced(namespace string) (string, error) {
	return os.Getenv("GITT_ROOT_" + strings.ToUpper(namespace)), nil
}
