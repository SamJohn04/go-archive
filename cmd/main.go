package main

import (
	"fmt"
	"os"

	"github.com/SamJohn04/go-archive/internal"
)

func main() {
	result, err := internal.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Error while parsing the input:", err)
		return
	}

	err = internal.ArchiveIt(result.Source, result.Destination)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	if result.DeleteOriginal {
		err = os.RemoveAll(result.Source)
		if err != nil {
			fmt.Printf("Error while deleting %v: %v\n", result.Source, err)
		}
	}
}

