package internal

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
)

type Argument struct {
	Source         string
	Destination    string
	DeleteOriginal bool
}

func Parse() (Argument, error) {
	var source, destination string
	var deleteOriginal bool

	flag.BoolVar(&deleteOriginal, "d", false, "delete the original folder after archive")
	flag.StringVar(&destination, "o", "", "the output file (add a .zip to folder by default)")

	flag.Parse()

	if len(flag.Args()) < 1 {
		return Argument{}, errors.New("no source string")
	}

	source, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		return Argument{}, fmt.Errorf("while parsing the source string: %w", err)
	}

	if len(destination) == 0 {
		destination = source + ".zip"
	} else {
		destination, err = filepath.Abs(destination)
		if err != nil {
			return Argument{}, fmt.Errorf("while parsing the destination string: %w", err)
		}
	}

	return Argument{
		Source:         source,
		Destination:    destination,
		DeleteOriginal: deleteOriginal,
	}, nil
}
