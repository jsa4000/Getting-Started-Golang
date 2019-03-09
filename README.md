# Getting-Started-Golang
Repository to learn the basics of the Go language

## Install

- Download the Go binaris from the golang [website](https://golang.org/dl/)
- Extract the content into a local folder

        tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz

- Set the environment variable to include the go binaries

        vi $HOME/.profile:
        export PATH=$PATH:/usr/local/go/bin

- Force to update the new variables

        source $HOME/.profile

- Check the version installed

        go version

  > $GOROOT variable is not necessary, just add the go executable into the binary path and set $GOPATH

- Tools to install (by vscode)
  - gocode
  - gopkgs
  - go-outline
  - go-symbols
  - guru
  - gorename
  - dlv
  - gocode-gomod
  - godef
  - goreturns
  - golint


## Hello World

- Create your workspace directory, `$HOME/go`. (If you'd like to use a different directory, you will need to set the `GOPATH` environment variable.)

  - Edit your `~/.bash_profile` to add the following line:

            export GOPATH=$HOME/go

  - Save and exit your editor. Then, source your `~/.bash_profile`.

            source ~/.bash_profile

- Next, make the directory `~/go/src/github.com/src/helloworld` inside your workspace, and in that directory create a file named `main.go` that looks like:

        ```go
        package main

        import "fmt"

        func main() {
            fmt.Printf("hello, world\n")
        }
        ```

- Then build it with the go tool:

        cd $HOME/go/src/github.com/src/helloworld
        go build

- The command above will build an executable named hello in the directory alongside your source code. Execute it to see the greeting:

        ./helloworld

- Or just run the file

        go run main.go

        # If there is more than one file in the folder user `*`
        go run *.fo
