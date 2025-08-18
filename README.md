# Go Archive

Archive your files on the Go.

## Flags

Use `-d` flag for deleting the original file.  
Use `-o FILENAME` flag to deciding the output file. By default, it has the same name as the parent, with .zip at the end.

The `-d` and `-o` tags are optional, but must be given before the source file.  

```cmd
archive [-d] [-o FILENAME] SOURCE_FILE_OR_DIR
```

Side note: the README assumes you have run `go build -o archive cmd/main.go` and moved the file appropriately.
