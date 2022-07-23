# Go project template

This repository is a project template for Go language for mainly backend server. Directory structure is following.

```
.
├── go.mod
├── go.sum
└── pkg
    ├── controller
    │   ├── cmd
    │   │   └── cli.go
    │   └── server
    │       └── server.go
    ├── infra
    │   ├── clients.go
    │   └── db
    │       └── client.go
    ├── model
    │   ├── context.go
    │   ├── errors.go
    │   ├── message.go
    │   └── user.go
    ├── usecase
    │   └── usecase.go
    └── utils
        └── logger.go
```
