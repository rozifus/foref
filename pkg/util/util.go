package util

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint //
func PrettyPrint(v interface{}) (err error) {
	s, err := PrettyString(v)
	if err != nil {
		return err
	}

	fmt.Println("%v", s)
	return nil
}

// PrettyString //
func PrettyString(v interface{}) (s string, err error) {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// CombineErrors //
func CombineErrors(errs []error) (error, error) {
	if len(errs) < 1 {
		return nil, fmt.Errorf("No errors to combine")
	}

	s, err := PrettyString(errs)
	if err != nil {
		return nil, fmt.Errorf("Could not combine errors, info omitted '%v'", err)
	}

	return fmt.Errorf("%s", s), nil
}

func PrintErrors(errs []error) {
	for err := range errs {
		fmt.Printf("%v", err)
	}
}
