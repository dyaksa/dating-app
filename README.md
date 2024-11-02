# Dating App

Simple Dating App Application

## Table of Contents

- [Getting Started](#getting-started)
- [Directory Structure](#directory-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)

## Getting Started

This service Built in Go, it is structured for modularity, maintainability, and scalability.

## Directory Structure

```plaintext
/project-root
├── cmd/
│   └── api
│       └── api.go        # API initialization and configuration
│   └── main.go           # Main entry point for the service
├── internal/
│   ├── app/              # Core application logic
│   ├── config/           # Configuration management
│   ├── dto/              # Data transfer objects
│   ├── infra/            # Infrastructure setup (database, external services)
│   ├── modules/          # Modularized business logic
│   ├── server/           # Server setup and routing
│   └── utils/            # Utility functions
└── pkg/                  # Shared packages and libraries
```

## Installation

Clone The Repository

```bash
git clone https://github.com/dyaksa/dating-app.git
```

Navigate to the project directory:

```bash
cd dating-app
```

Install dependencies:

```bash
go mod tidy
```

## Usage

Install Goose For Migrate DB

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

This will install the goose binary to your $GOPATH/bin directory.

Binary too big? Build a lite version by excluding the drivers you don't need:

```bash
go build -tags='no_postgres no_mysql no_sqlite3 no_ydb' -o goose ./cmd/goose

# Available build tags:
#   no_clickhouse  no_libsql   no_mssql    no_mysql
#   no_postgres    no_sqlite3  no_vertica  no_ydb
```

For macOS users goose is available as a Homebrew

```bash
brew install goose
```

To migrate database migration, use:

```bash
make migrate
```

To start the service, use:

```bash
make run
```

## Configuration

Config env configuration before running :

```bash
APP_PORT=

DB_HOST=
DB_PORT=
DB_USER=
DB_NAME=
DB_PASSWORD=

GIN_MODE=

SECRET_KEY=
```
