# Soccer API

[Soccer API](http:http://localhost:8070/) default port

## Setup Local Environment

### Prerequisites

+ Git
+ Go (minimum 1.16 or above, make sure you set your local machine environment with `GO111MODULE=on`)
+ [Mockgen](https://github.com/golang/mock)

### Installation

+ Clone this repository.

+ Make the environment files. Adjust your local configuration.
```bash
$ cp config.yaml.tpl config.yaml
```

+ Run the app. The app will run inside the local machine with exposed port configured in the env (by default: [localhost:8070](http://localhost:8070))
```bash
$ go run *main*.go
```

#### Mock an interface
Use `mockgen` to create a new mock. Tips: place mock file inside [mocks](mocks) directory
```bash
$ mockgen -package=mock -source=/path/to/interface/file -destination=/path/to/generated/mock/file
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

### API

Endpoint in this API
```
/v1/team [GET] -> get team with player
/v1/team [POST] -> store team

/v1/player [GET] -> get player
/v1/player [POST] -> store player
```

Request body store team
```json
{
    "name": "Furano"
}
```

Request body store player
```json
{
    "team_name": "Meiwa",
    "name": "Kagawa"
}
```

