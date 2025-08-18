package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/SamJohn04/go-archive/internal"
)

func main() {
	result, err := internal.Parse()
	if err != nil {
		fmt.Println("Error while parsing the input:", err)
		return
	}

	err = archiveIt(result.Source, result.Destination)
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

func archiveIt(source, destination string) error {
	if _, err := os.Stat(destination); !errors.Is(err, os.ErrNotExist) {
		return errors.New("zip file already exists")
	}
	srcStat, err := os.Stat(source)
	if errors.Is(err, os.ErrNotExist) {
		return errors.New("source file does not exist")
	} else if err != nil {
		return fmt.Errorf("something went wrong while accessing file: %w", err)
	}

	zipFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	if !srcStat.IsDir() {
		err = archiveFile(source, filepath.Base(source), srcStat, archive)
		if err != nil {
			return err
		}

		if err = archive.Flush(); err != nil {
			return err
		}
		return nil
	}

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

		return archiveFile(path, path[len(source)+1:], info, archive)
	})
	if err != nil {
		return err
	}

	if err := archive.Flush(); err != nil {
		return err
	}
	return nil
}

func archiveFile(sourceFilePath, archiveFilePath string, info fs.FileInfo, destinationFile *zip.Writer) error {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = archiveFilePath
	header.Method = zip.Deflate

	writer, err := destinationFile.CreateHeader(header)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	file, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(writer, file)

	return err
}
