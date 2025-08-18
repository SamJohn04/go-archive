# Go Archive

Archive your files on the Go.

## Initialization

Use `go build -o archive cmd/main.go` to create archive.  
Once created, move it to a directory that is accessible from the whole system.  

On the other hand, you can also use `go run cmd/main.go ...`.

## Flags

Use `-d` flag for deleting the original file.  
Use `-o FILENAME` flag to deciding the output file. By default, it has the same name as the parent, with .zip at the end.

The `-d` and `-o` tags are optional, but must be given before the source file.  

```cmd
archive [-d] [-o FILENAME] SOURCE_FILE_OR_DIR
```
