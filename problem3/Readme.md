## Setup Local Environment

### Prerequisites

+ Git
+ Go (minimum 1.16 or above, make sure you set your local machine environment with `GO111MODULE=on`)
+ [Mockgen](https://github.com/golang/mock)

### Installation

+ Clone this repository.

+ Run the app. The app will run inside the by default it contain 20 cake and 25 apple
```bash
$ go run answer.go
```

+ You can modify cake and apple by adding the flag
```bash
$ go run answer.go -cake=12 -apple=3
```

#### Add dependency
Use `go mod` as dependency tool.
```bash
$ go get ./...
$ go mod tidy
```

#### Unit test
Use `go test` to run unit test.
```bash
$ go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
```
