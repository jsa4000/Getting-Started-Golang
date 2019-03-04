# Swagger

## Go-Swagger

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

## SwaggerUI

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

[Go-Swagger](https://github.com/go-swagger/go-swagger)