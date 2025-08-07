package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

type argument struct {
	source      string
	destination string
}

func parse(cmdArgs []string) (argument, error) {
	var source, destination string

	if len(cmdArgs) < 1 {
		return argument{}, errors.New("no source string")
	}

	source, err := filepath.Abs(cmdArgs[0])
	if err != nil {
		return argument{}, fmt.Errorf("while parsing the source string: %w", err)
	}

	if len(cmdArgs) == 1 {
		destination = source + ".zip"
	} else {
		destination, err = filepath.Abs(cmdArgs[1])
		if err != nil {
			return argument{}, fmt.Errorf("while parsing the destination string: %w", err)
		}
	}

	return argument{
		source:      source,
		destination: destination,
	}, nil
}
