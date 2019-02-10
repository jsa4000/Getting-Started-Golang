# WebApp

## Installation

Install Go binaries (+v1.11.x)

Install the dependencies

    go install

Run the server (configuration by default)

    go run main.go

To generate the binary use ``go build``

## Testing

Go testing look for every files with ``'_test.go'`` at the end.
Each test must start with ``'Test..'``

To run the tests

    # Run all the test inside the folder
    go test

    # Specific file (without dependencies..)
    go test module_test.go

    # Run a specific a test
    go test -run TestSomething

    # Verbose option to see the tests performed
    go test -v ./...

    # Verbose option to see the tests performed
    go test -v

## Configuration

The configuration file ``config.yaml`` is the following

```yaml
app:
  name: WebApp
  
logging:
  level: debug

server:
  port: 8080
  writeTimeout: 15
  readTimeout: 15
  idleTimeout: 60

repository:
  roles:
    mocked: enabled
    provider: mongodb
    mongodb:
      database: roles
      url: mongodb://db1.example.net:27017,db2.example.net:2500/?replicaSet=test
  users:
    mocked: enabled
    provider: mongodb
    mongodb:
      database: users
      url: mongodb://db1.example.net:27017,db2.example.net:2500/?replicaSet=test
```

## Docker

- Build docker image

      docker build -t webapp-go .

- Execute docker container (copy configuration file)

      docker run -p 8080:8080 -v //d/DEVELOPMENT/Github/Getting-Started-Golang/src/webapp/config.docker.yaml:/config.yaml -t webapp-go

      docker run -p 8080:8080 -v $(pwd)/config.docker.yaml:/config.yaml -t webapp-go

- To use the same network as configured docker-compose ``webapp-network``

      docker network ls

      docker run --network=webapp-network -p 8080:8080 -v //d/DEVELOPMENT/Github/Getting-Started-Golang/src/webapp/config.docker.yaml:/config.yaml -t webapp-go

- Test Server

    http://localhost:8080

    http://dockerhost:8080/users

- Docker Shutdown (gracefully shutdown)

      # Get the PID of the container to stop
      docker ps

      # Get logs from another shell
      docker logs <container-id> -f

      # Stop the server (vs kill)
      docker stop <container-id>

## Profiling

### Graphviz - Graph Visualization Software

Install ``Graphviz`` for the current platform to be able to see ``svg`` images and to use the following tools.

Add ``Graphviz/bin`` folder into $PATH environment variable

    /usr/local/Graphviz/bin

### pprof 

The first step to profiling a Go program is to enable ``profiling``. Support for profiling benchmarks built with the standard testing package is built into ``go test``. For example, the following command runs benchmarks in the current directory and writes the CPU and memory profiles to ``cpu.prof`` and ``mem.prof``:

    go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

### http/pprof 

To use ``pprof``, link this package into your program:

    import _ "net/http/pprof"

Then use the pprof tool to look at the heap profile:

> It must be used ``jmeter``, ``go-wrk``, etc perform some calls to the application.

    go tool pprof http://localhost:8080/debug/pprof/heap

    # seconds < server.writeTimeout
    go tool pprof http://localhost:8080/debug/pprof/profile?seconds=29

    #The input 'web' to generate the svg file
    (pprof) web

Profiles can then be visualized with the ``pprof`` tool:

    go tool pprof cpu.prof

### go-torch

    go-torch --seconds 5 http://localhost:8080/debug/pprof/profile

### wrk

wrk](https://github.com/wg/wrk)

This runs a benchmark for 30 seconds, using 12 threads, and keeping 400 HTTP connections open.

    wrk -t12 -c400 -d30s http://127.0.0.1:8080/index.html

## Swagger

### Go-Swagger

**API First design** is a very standard good practice, since it forces to start the definition of the endpoints. 
Standard ways to define Rest API is using Open API and Swagger definitions.

[Go-Swagger](https://github.com/go-swagger/go-swagger)

First, creates a definition using swagger. (*VSCode Swagger plugin*)

Then generates the server, client, stub, etc from previous definition. The official ``swagger-codegen`` library does not implements all the functionality defined in the definitions, suh as authentication, middleware, validations, types, etc..

    docker pull quay.io/goswagger/swagger

    docker run -it quay.io/goswagger/swagger version

Generate the json specification from the yaml

    docker run -it -w /go/src -v //d/DEVELOPMENT/Github/Getting-Started-Golang/src/webapp/files:/go/src quay.io/goswagger/swagger generate spec -i swagger.yaml -o swagger.json

Generate server code from previous specification

    docker run -it -w /go/src -v //d/DEVELOPMENT/Github/Getting-Started-Golang/src/webapp/files:/go/src quay.io/goswagger/swagger generate server -f swagger.json

Generate client code from previous specification

    docker run -it -w /go/src -v //d/DEVELOPMENT/Github/Getting-Started-Golang/src/webapp/files:/go/src quay.io/goswagger/swagger generate client -f swagger.json

### SwaggerUI

- Downloading SwaggerUI files

  SwaggerUI can be downloaded from their [GitHub Repo](https://github.com/swagger-api/swagger-ui). 
  Once downloaded, place the content of ``dist`` folder somewhere in your Go project. For example, static/swaggerui.

  After that, move swagger.json file to swaggerui folder, and inside index.html change url to ./swagger.json (url: "./swagger.json").

- Serve using net/http

    ```go
    fs := http.FileServer(http.Dir("./swaggerui"))
    http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
    ```

- Serve using Gorilla Mux (commit)

    ```go
    sh := http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/")))
    r.PathPrefix("/swaggerui/").Handler(sh)
    ```

## References

### Frameworks

- [Go-Kit](https://github.com/asaskevich/govalidator)
- [Go Gin](https://github.com/asaskevich/govalidator)

### Validations

- [Go Validator](https://github.com/asaskevich/govalidator)
- [Go-Playground validator](https://github.com/go-playground/validator)