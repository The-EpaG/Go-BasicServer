# backend-starter-kit-go

The **backend-starter-kit-go** is a starter kit for Go projects that enables you to easily kickstart your Golang project. This project comes pre-configured with MySQL support and provides the necessary tools to streamline your development process.

## Installation

To install and get started with the **backend-starter-kit-go**, follow these steps:

### Clone the repository:

```sh
git clone git@github.com/The-EpaG/Go-BasicServer
```

### Change to the project directory:

```sh
cd backend-starter-kit-go
```


### Create containers:

```sh
cd docker
cp .env.local .env
docker compose up -d
```

### Quick setup

before running any command, make sure you have the following prerequisites installed:
-  Go 1.23.4
-  Make >= 4.3

```sh
go mod tidy
make tools
```
this will install the required dependencies and the following dev tools in the `./bin` directory:
- [CompileDaemon](github.com/githubnemo/CompileDaemon)
- gotestsum
- golangci-lint

it will also install the git `pre-commit` and `prepare-commit-msg` hooks

then create the secret.yaml config file with the following command:

```sh
cp config/secret.yaml.local config/secret.yaml
```

open the `config/secret.yaml` file and provide your data and database credentials.

## Starting the Project

to start the project, while watching for changes, run:
```sh
make watch
# or just 'make' since the 'watch' target is the default
```

or just build the project standalone executable:
```sh
make build
```

you can run the server using the standalone executable with:
```sh
./main serve
```

the env target is specified in the `config/secret.yaml` file in the `app.env` field, for a production environment, set it to `prod`.

### Run tests
```sh
make test
```
this runs all tests types (unit, integration, e2e)

to only run tests of a specific type, use the specific target:
```sh
make test-e2e
```

