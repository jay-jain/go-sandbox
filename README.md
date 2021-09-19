# Go Sandbox
## Install

### Version
```
go version
```

### Upgrade

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

## References
https://gobyexample.com