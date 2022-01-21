# Fibonacci RESTful API

## Overview
HTTP server that listens on a given port following Hexagonal architecture (*ports & adapters*)
It supports multiple connections simultaneously, and provides the following endpoints:

 - `/fib/algorithm`
    - input: form field named 'input' using POST to provide the value n of which the nth fibonacci number will be calculated. Input must be between 1 and 99999. POST
    - input: algorithm with which to calculate the fin number. Options: *math*, *recursive*, *iterate*
    - output: An incrementing identifier returns immediately.
    - example_input: curl --data "input=50" http://localhost:8000/fib/math
    - example_outut: 1
 
  - `/find/id`
    - input: id which was passed to the user from the /fib/algorithm endpoint. GET
    - output: json encoded list of details about the request. If the number is not done calculating "status" will be set to incomplete. Duration is in microseconds.
    - example_input: curl http://localhost:8000/find/1
    - example_output: {"Input":50,"Fib":7778742049,"Duration":123,"Algo":"math","Status":"complete","Id":1}
 
 - `/shutdown`
    - input: None. 
    - output: Server will gracefully shutdown after waiting for all active requests to complete.
    

## Installation
### Setup

After installing GO, Clone this repo and launch server in a terminal with go run main.go port#
  - ex: go run main.go 8000. The server automatically starts on localhost

In a new terminal send your POST and GET requests.

## Architecture

Code architecture follows hexagonal architecture principles, also known as *ports and adapters*.

This architecture is divided in three main layers:

- **Application**:  The outer layer. Handlers and all I/O related stuff (web framework, DB, ...). Anything that can change by an "external" cause (not by your decision), is in this layer. 

- **Service**: Use cases. Actions triggered by API calls, represented by application services. It includes repositories specific interfaces, known as *adapters*.

- **Domain**: Inner layer. Business logic and rules goes here. Repositories Interfaces, known as *ports*, belongs to this layer.

I have also included an data transfer object (dto) abstraction layer. This allows us to control exactly what is passed back to the client.

## Testing 

Tests have been built with golangs built in testing package. In root directory of repo run `go test ./...`

## Considerations made

- I considered several approaches to the routing including gorilla mux and some regex strategies, but in the end I thought a "no router" approach was the most maintainable.  
- I did not hook a database up to this server as I didn't want to over engineer the prompt, but due the hexagonal architecture it would be pretty trivial to add one. 
- I took liberties with the prompt, making the server more RESTful.
- I only included tests for the business logic (domain) in this example server but I would usually have tests for app, domain, and service sides.
- If I was to iterate on this I would add more debug logging functionality, and more error checking. 
