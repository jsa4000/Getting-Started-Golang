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

The configuration file ``webapp.yaml`` is the following

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

      docker run -p 8080:8080 -v //d/DEVELOPMENT/Github/Getting-Started-Golang/src/webapp/webapp.yaml:/webapp.yaml -t webapp-go

      docker run -p 8080:8080 -v config/webapp.yaml:/webapp.yaml -t webapp-go

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

## pprof 

The first step to profiling a Go program is to enable ``profiling``. Support for profiling benchmarks built with the standard testing package is built into ``go test``. For example, the following command runs benchmarks in the current directory and writes the CPU and memory profiles to ``cpu.prof`` and ``mem.prof``:

    go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

## http/pprof 

 To use ``pprof``, link this package into your program:

    import _ "net/http/pprof"

Then use the pprof tool to look at the heap profile: 

    go tool pprof http://localhost:8080/debug/pprof/heap

## References

### Frameworks

- [Go-Kit](https://github.com/asaskevich/govalidator)
- [Go Gin](https://github.com/asaskevich/govalidator)

### Validations

- [Go Validator](https://github.com/asaskevich/govalidator)
- [Go-Playground validator](https://github.com/go-playground/validator)