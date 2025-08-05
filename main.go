package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
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

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
	if err != nil {
		return err
	}

	if err := archive.Flush(); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Expecting source and target")
		return
	}

	source := os.Args[1]
	target := os.Args[2]

	err := archiveIt(source, target)
	if err != nil {
		fmt.Println("Something went wrong: ", err)
	}
}
