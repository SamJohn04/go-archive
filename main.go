package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

	base := filepath.Base(source)

	err = filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if path == source {
				return nil
			}
			path += "/"
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = path[len(base)+1:]
		header.Method = zip.Deflate

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		return nil
	})

	return nil
}

func main() {
	fmt.Println("WIP")
}
