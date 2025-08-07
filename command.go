package main

import "errors"

type argument struct {
	source      string
	destination string
}

func parse(cmdArgs []string) (argument, error) {
	if len(cmdArgs) < 1 {
		return argument{}, errors.New("no source string")
	}

	if len(cmdArgs) < 2 {
		return argument{}, errors.New("no destination string")
	}

	return argument{
		source:      cmdArgs[0],
		destination: cmdArgs[1],
	}, nil
}
