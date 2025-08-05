package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"os"
)

func archiveIt(source, target string) error {
	if _, err := os.Stat(target); !errors.Is(err, os.ErrNotExist) {
		return errors.New("file already exists")
	}

	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	return nil
}

func main() {
	fmt.Println("WIP")
}
