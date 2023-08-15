# ProSpace Backend Developer Test
Built with Go

## How to install and run
Make sure a version of Go is installed on your machine
- Clone the repository
- In your preferred terminal, run
```bash
go mod tidy
```
- To run, in your preffered terminal, run
```bash
go run main.go
```

## Architecture
[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) 

## Structure

```bash
├── repository
│   ├── repository_adapter.go
│   ├── repository_test.go
│   ├── repository.go
│   └── mock
│       └── mock_repository.go
├── helpers
│   ├── indexOf_test.go
│   └── indexOf.go
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```