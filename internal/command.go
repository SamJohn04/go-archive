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

var deleteOriginal = flag.CommandLine.Bool("d", false, "delete the original folder after archive")
var destination = flag.CommandLine.String("o", "", "the output file (adds a .zip to source by default)")

func Parse(arguments []string) (Argument, error) {
	var source string

	err := flag.CommandLine.Parse(arguments)
	if err != nil {
		resetArgsToDefault()
		return Argument{}, err
	}

	if len(flag.Args()) < 1 {
		resetArgsToDefault()
		return Argument{}, errors.New("no source string")
	}

	source, err = filepath.Abs(flag.Arg(0))
	if err != nil {
		resetArgsToDefault()
		return Argument{}, fmt.Errorf("while parsing the source string: %w", err)
	}

	if len(*destination) == 0 {
		*destination = source + ".zip"
	} else {
		*destination, err = filepath.Abs(*destination)
		if err != nil {
			resetArgsToDefault()
			return Argument{}, fmt.Errorf("while parsing the destination string: %w", err)
		}
	}

	argument := Argument{
		Source:         source,
		Destination:    *destination,
		DeleteOriginal: *deleteOriginal,
	}
	resetArgsToDefault()
	return argument, nil
}

func resetArgsToDefault() {
	*deleteOriginal = false;
	*destination = "";
}
