# Bashctl

## Initializing

Firstly, the module must be inititialized (go > 1.11.x)

    go mod init modulename

This creates a new file called `go.mod` with all the information needed to get all the packages.

## Packages

### Logging

The library used is `logrus`, since it has a lot of community.

```go
import (
    log "github.com/sirupsen/logrus"
)
```

There are several ways to format the logs (`default`, `json`, etc..):

- Default Formatter

    ```go
    // Default formatter
    log.SetFormatter(&log.TextFormatter{})
    log.Info("Something noteworthy happened!")
    log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info("My first ssl event from golang")
    ```

- JSON Formatter

    ```go
    // JSON formatter
    log.SetFormatter(&log.JSONFormatter{})
    log.Info("Something noteworthy happened!")
    log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info("My first ssl event from golang")
    ```

## Command-line flags

Command-line flags are a common way to specify options for **command-line** programs. For example, in `wc -l` the `-l` is a *command-line flag*.

Go provides a `flag` package supporting basic command-line flag parsing.

```go
import "flag"
```

Basic flag declarations are available for `string`, `integer`, and `boolean` options. Here we declare a string flag word with a default value "foo" and a short description. This `flag.String` function returns a string *pointer* (**not a string value**).

    // Definition flags.String
    func String(name string, value string, usage string) *string

```go
cmd := flag.String("cmd", "ls", "List all the files in a folder.")
flag.StringVar(&cmd, "c", "ls", "List all the files in a folder.")
```

Once all flags are declared, call `flag.Parse()` to execute the command-line parsing.

```go
flag.Parse()
fmt.Println("command:", *command)
```

User `flag.Args()` to get all the other parameters that are considered non-flags.

## Usage

Get the help with the available options

    go run main.go -h

```txt
Usage of /tmp/go-build083531499/b001/exe/main:
  -a string
        Arguments to execute with the command (shorthand)
  -args string
        Arguments to execute with the command
  -c string
        Command to execute (shorthand)
  -cmd string
        Command to execute
  -v    View all the logs and traces (shorthand)
  -verbose
        View all the logs and traces
```

Execute the command using one of the following ways

    go run main.go -cmd ls
    go run main.go --cmd ls
    go run main.go -c=ls

    // Following usage will thrown an error
    go run main.go -c  

As advanced example use the following command

    go run main.go -c ls -v=true --a=-ls,/etc

Build the package and run the command, similar of using to `go run main.go`

    # Generates the binary 'bashctl'
    go build

   ./bashctl -c ls -v=true --a=-ls,/etc

## References

- [Testify (testing)](https://github.com/stretchr/testify)
- Kubernetes (container orchestration)
- Ginkgo (testing)
- [Logrus (logging)](https://github.com/sirupsen/logrus)
- Gomega (testing)
- [glog (logging)](https://github.com/golang/glog)
- gocheck (testing)
- AWS SDK (cloud tools)
- errors (error handling)
- cobra (productivity)