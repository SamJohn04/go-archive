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
		header.Name = path[len(source)+1:]
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
	result, err := Parse()
	if err != nil {
		fmt.Println("Error while parsing the input:", err)
		return
	}

	err = archiveIt(result.Source, result.Destination)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	if !result.DeleteOriginal {
		return
	}

	err = os.RemoveAll(result.Source)
	if err != nil {
		fmt.Printf("Error while deleting %v: %v\n", result.Source, err)
	}
}
