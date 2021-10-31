# Go Sandbox
## Install

### Version
```
go version
```

### Install/ Upgrade

- Uninstall existing version:

```
sudo rm -rf /usr/local/go
```

- Download new version from [here](https://golang.org/dl/)

- Extract the archive

```
$ sudo tar -C /usr/local -xzf /home/nikhita/Downloads/go<VERSION>.linux-amd64.tar.gz
```
- Validate your `PATH` variable

```
echo $PATH | grep "/usr/local/go/bin"
```

## Build / Run

### Build binary
```
go build *.go
```

A binary should exist with the same name as the `.go` file that was used for the build above.

### Run
```
go run hello-world.go
```

### Install
go install

### Get
Resolves command-line arguments to packages at specific module versions,
updates go.mod to require those versions, downloads source code into the
module cache, then builds and installs the named packages.

```
go get example.com/pkg@v1.2.3
```

### Clean
Removes object files from package source directories.

```
go clean
```

## Modules
A module is a collection of related Go packages that are versioned together as a single unit.
Modules record precise dependency requirements and create reproducible builds. 

### Initialize a module
```
go mod init <MODULE_NAME>
```

This creates a `go.mod` file.

### `GOPATH`
The Go path is used to resolve import statements.

The GOPATH environment variable lists places to look for Go code.

If the environment variable is unset, GOPATH defaults to a subdirectory named "go" in the user's home directory ($HOME/go on Unix, %USERPROFILE%\go on Windows),unless that directory holds a Go distribution.

Run `go env GOPATH` to see the current GOPATH.

`GOPATH` must not be the same path as your Go installation

## References
https://gobyexample.com