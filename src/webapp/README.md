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

    go test

    # Specific file (without dependencies..)
    go test module_test.go

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