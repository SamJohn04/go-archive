package main

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
)

type argument struct {
	source         string
	destination    string
	deleteOriginal bool
}

func parse() (argument, error) {
	var source, destination string
	var deleteOriginal bool

	flag.BoolVar(&deleteOriginal, "d", false, "delete the original folder after archive")
	flag.StringVar(&destination, "o", "", "the output file (add a .zip to folder by default)")

	flag.Parse()

	if len(flag.Args()) < 1 {
		return argument{}, errors.New("no source string")
	}

	source, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		return argument{}, fmt.Errorf("while parsing the source string: %w", err)
	}

	if len(destination) == 0 {
		destination = source + ".zip"
	} else {
		destination, err = filepath.Abs(destination)
		if err != nil {
			return argument{}, fmt.Errorf("while parsing the destination string: %w", err)
		}
	}

	return argument{
		source:         source,
		destination:    destination,
		deleteOriginal: deleteOriginal,
	}, nil
}
